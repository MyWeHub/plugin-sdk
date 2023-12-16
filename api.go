package wehublib

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/MyWeHub/plugin_sdk/gen/pluginrunner"
	"github.com/MyWeHub/plugin_sdk/util"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var pluginName = util.GetEnv("PLUGIN_NAME", false, "", true)

func (s *grpcServer) RunTestv2(ctx context.Context, input *pb.InputTestRequestV2) (*pb.InputTestResponseV2, error) {
	if input.Inputs == nil {
		return nil, errors.New("input is empty")
	}

	workflowData := ""
	if input.NodeId != nil {
		workflowData = *input.NodeId
	}

	conf, err := decodeConf(input.Configuration, s.nats.ConfigType)
	if err != nil {
		return nil, err
	}

	vpb, err := s.service.Process(ctx, input.Inputs, conf, input.Action, workflowData) // TODO: check NodeID
	if err != nil {
		return nil, err
	}

	return &pb.InputTestResponseV2{Outputs: vpb}, nil
}

func (s *grpcServer) RunV2(ctx context.Context, input *pb.InputRequestV2) (*pb.InputTestResponseV2, error) {
	_, span := tracer.Start(ctx, fmt.Sprintf("Run %s", pluginName))
	span.SetAttributes(attribute.String("transformationId", input.TransformationId))
	span.SetAttributes(attribute.String("nodeId", input.NodeId))
	defer span.End()
	if node, ok := s.nats.Cache[input.NodeId]; !ok {
		err := &WrongConfigurationError{fmt.Sprintf("Configuration missing for node %s", input.NodeId)}
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)
		return nil, status.Convert(err).Err()
	} else {
		vpb, err := s.service.Process(ctx, input.Inputs, node.Configuration, input.Action, input.NodeId)
		if err != nil {
			return nil, err
		}

		return &pb.InputTestResponseV2{Outputs: vpb}, nil
	}
}

func decodeConf(b []byte, config proto.Message) (proto.Message, error) {
	newConfig := proto.Clone(config)
	err := protojson.Unmarshal(b, newConfig)
	if err != nil {
		return nil, err
	}

	return newConfig, nil
}

type WrongConfigurationError struct {
	Msg string
}

func (e *WrongConfigurationError) Error() string {
	return e.Msg
}

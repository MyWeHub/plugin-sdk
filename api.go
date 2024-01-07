package wehublib

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/MyWeHub/plugin-sdk/gen/pluginrunner"
	"github.com/MyWeHub/plugin-sdk/util"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var pluginName = util.GetEnv("PLUGIN_NAME", false, "", true)

func (s *grpcServer) RunTestv2(ctx context.Context, input *pb.InputTestRequestV2) (*pb.InputTestResponseV2, error) {
	if input.Inputs == nil {
		return nil, status.Convert(errors.New("input is empty")).Err()
	}

	workflowData := ""
	if input.NodeId != nil {
		workflowData = *input.NodeId
	}

	conf, err := decodeConf(input.Configuration, s.configType)
	if err != nil {
		logger.Error("decodeConf", zap.Error(err))
		return nil, status.Convert(err).Err()
	}

	out, err := s.service.Process(ctx, input.Inputs, conf, input.Action, workflowData) // TODO: check NodeID
	if err != nil {
		logger.Error("Process", zap.Error(err))
		return nil, status.Convert(err).Err()
	}

	return out, nil
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
		out, err := s.service.Process(ctx, input.Inputs, node.Configuration, input.Action, input.NodeId)
		if err != nil {
			logger.Error("Process", zap.Error(err))
			return nil, status.Convert(err).Err()
		}

		return out, nil
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

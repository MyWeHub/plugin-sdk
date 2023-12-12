package wehublib

import (
	"context"
	pb "dev.azure.com/WeConnectTechnology/ExchangeHub/_git/wehublib.git/gen/pluginrunner"
	"dev.azure.com/WeConnectTechnology/ExchangeHub/_git/wehublib.git/util"
	"errors"
	"fmt"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/status"
)

var pluginName = util.GetEnv("PLUGIN_NAME", false, "", true)

func (s *grpcServer) RunTestv2(ctx context.Context, input *pb.InputTestRequestV2) (*pb.InputTestResponseV2, error) {
	if input.Inputs == nil {
		return nil, errors.New("input is empty")
	}

	/*var config pbconf.Configuration
	err := protojson.Unmarshal(input.Configuration, &config)
	if err != nil {
		return nil, fmt.Errorf("wrong configuration: %w", err)
	}*/

	vpb, err := s.service.Process(ctx, input.Inputs, nil, input.Action, *input.NodeId)
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
		vpb, err := s.service.Process(ctx, input.Inputs, node, input.Action, input.NodeId)
		if err != nil {
			return nil, err
		}

		return &pb.InputTestResponseV2{Outputs: vpb}, nil
	}
}

type WrongConfigurationError struct {
	Msg string
}

func (e *WrongConfigurationError) Error() string {
	return e.Msg
}

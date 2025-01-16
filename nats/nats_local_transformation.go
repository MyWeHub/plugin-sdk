//go:build local_tranformation
// +build local_tranformation

package nats

import (
	"context"
	"flag"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io"
	"os"
)

// go build -tags local_transformation -o bin_name .

var (
	logger *zap.Logger
	tracer trace.Tracer
)

func SetTelemetry(l *zap.Logger, t trace.Tracer) {
	logger = l
	tracer = t
}

type Nats struct {
	Cache      map[string]*NodeConfig
	ConfigType proto.Message
}

type ListenerOptions struct {
	UpdateFunc func(nc *NodeConfig)
	RemoveFunc func(nc *NodeConfig)
}

func New(configType proto.Message) *Nats {
	configAddr := flag.String("config", "", "Configuration JSON file relative address")
	workflowID := flag.String("workflowID", "", "Workflow Version ID")

	// Parse the flags
	flag.Parse()

	if *configAddr == "" {
		logger.Fatal("Must specify -config flag")
	}

	if *workflowID == "" {
		logger.Fatal("Must specify -workflowID flag")
	}

	// Open the JSON file
	file, err := os.Open(*configAddr)
	if err != nil {
		logger.Fatal("Failed to open JSON file", zap.Error(err))
	}
	defer file.Close()

	// Read the file content
	data, err := io.ReadAll(file)
	if err != nil {
		logger.Fatal("Failed to read JSON file", zap.Error(err))
	}

	// Unmarshal the JSON into the Protobuf message
	if err = protojson.Unmarshal(data, configType); err != nil {
		logger.Fatal("Failed to unmarshal JSON to Protobuf", zap.Error(err))
	}

	n := &Nats{}

	// Add Config to cache
	n.Cache = make(map[string]*NodeConfig)
	n.Cache[*workflowID] = &NodeConfig{
		NodeType:      *workflowID,
		ID:            *workflowID,
		WorkflowID:    *workflowID,
		Configuration: configType,
	}

	return n
}

func (n *Nats) Listen(ctx context.Context, opts ...*ListenerOptions) {}

func (n *Nats) Close() {}

type NodeConfig struct {
	NodeType      string
	ID            string `bson:"_id"`
	WorkflowID    string
	Configuration proto.Message
	ClientID      string
}

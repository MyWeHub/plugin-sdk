package nats

import (
	"context"
	"dev.azure.com/WeConnectTechnology/ExchangeHub/_git/wehublib.git/telemetry"
	"dev.azure.com/WeConnectTechnology/ExchangeHub/_git/wehublib.git/util"
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"time"
)

var (
	logger *zap.Logger
	tracer trace.Tracer
)

type Nats struct {
	conn       *nats.EncodedConn
	Cache      map[string]*NodeConfig
	configType interface{}
}

func New(t *telemetry.Telemetry) *Nats {
	logger = t.GetLogger()
	tracer = t.GetTracer()

	natsURL := util.GetEnv("NATS", false, "nats://localhost:4222", false)
	conn, err := nats.Connect(natsURL)
	if err != nil {
		logger.Error("failed to connect to NATS", zap.String("nats URL", natsURL), zap.Error(err))
		panic(err)
	}

	encodedConn, err := nats.NewEncodedConn(conn, nats.JSON_ENCODER)
	if err != nil {
		logger.Error("failed to Encode Connection", zap.Error(err))
		panic(err)
	}

	return &Nats{
		conn:  encodedConn,
		Cache: make(map[string]*NodeConfig),
	}
}

func (n *Nats) updateCache(configs *[]NodeConfig) {
	for _, nodeConfig := range *configs {
		nc := nodeConfig
		n.Cache[nodeConfig.ID] = &nc
	}
}

func (n *Nats) Listen(ctx context.Context) {
	pluginName := util.GetEnv("PLUGIN_NAME", false, "", true)

	ctx, span := tracer.Start(ctx, "Request plugin configuration")
	defer span.End()

	req := fmt.Sprintf("requestPluginConfig.%s", pluginName)
	span.SetAttributes(attribute.String("topic", req))

	natsConfigs := make([]NodeConfig, 0)
	err := n.conn.Request(req, "", &natsConfigs, 10*time.Second)
	if err != nil {
		logger.Error("Request failed", zap.Error(err))
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)
		panic(err)
	}

	configs := make([]NodeConfig, 0, len(natsConfigs))
	for _, conf := range natsConfigs {
		configs = append(configs, conf)
	}
	logger.Info("Received plugin config", zap.Any("configs", configs))

	n.updateCache(&configs)

	n.conn.Subscribe(fmt.Sprintf("refresh.%s.*", pluginName), func(m *nats.Msg) {
		_, span := tracer.Start(ctx, "Received a config")
		err := json.Unmarshal(m.Data, &natsConfigs)
		if err != nil {
			zap.Error(err)
		}
		configs := make([]NodeConfig, 0, len(natsConfigs))
		for _, conf := range natsConfigs {
			configs = append(configs, conf)
		}
		defer span.End()
		logger.Info("Received a config", zap.Any("configs", configs))
		n.updateCache(&configs)
	})
}

func (n *Nats) Close() {
	n.conn.Close()
}

type nodeConfigNats struct {
	nodeType      string
	id            string `bson:"_id"`
	workflowID    string
	configuration json.RawMessage
	clientID      string
}

type NodeConfig struct {
	NodeType      string
	ID            string `bson:"_id"`
	WorkflowID    string
	Configuration interface{}
	ClientID      string
}

/*
func (s *NodeConfig) DecodeConfig(config proto.Message) error {
	err := protojson.Unmarshal(s.Configuration, config)
	if err != nil {
		return err
	}

	return nil
}*/

/*func (s *nodeConfigNats) decode(configType interface{}) (*NodeConfig, error) {
	reflect.TypeOf(configType)
	var config pbconf.Configuration
	err := protojson.Unmarshal(s.configuration, &config)
	if err != nil {
		return nil, err
	}

	return &NodeConfig{
		NodeType:      s.nodeType,
		ID:            s.id,
		WorkflowID:    s.workflowID,
		Configuration: &config,
		ClientID:      s.clientID,
	}, nil
}*/

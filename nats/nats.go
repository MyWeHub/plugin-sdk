package nats

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MyWeHub/plugin_sdk/util"
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"time"
)

var (
	logger *zap.Logger
	tracer trace.Tracer
)

func SetTelemetry(l *zap.Logger, t trace.Tracer) {
	logger = l
	tracer = t
}

type Nats struct {
	conn       *nats.EncodedConn
	Cache      map[string]*NodeConfig
	ConfigType proto.Message
}

func New(configType proto.Message) *Nats {
	if configType == nil {
		panic(errors.New("nats: ConfigType parameter is nil"))
	}

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
		conn:       encodedConn,
		Cache:      make(map[string]*NodeConfig),
		ConfigType: configType,
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

	natsConfigs := make([]NodeConfigNats, 0)
	err := n.conn.Request(req, "", &natsConfigs, 10*time.Second)
	if err != nil {
		logger.Error("Request failed", zap.Error(err))
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)
		panic(err)
	}

	configs := make([]NodeConfig, 0, len(natsConfigs))
	for _, conf := range natsConfigs {
		decodedConf, err := conf.decode(n.ConfigType)
		if err == nil {
			configs = append(configs, *decodedConf)
		} else {
			logger.Warn("decode nats config", zap.Error(err))
		}
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
			decodedConf, err := conf.decode(n.ConfigType)
			if err == nil {
				configs = append(configs, *decodedConf)
			} else {
				logger.Warn("decode nats config", zap.Error(err))
			}
		}
		defer span.End()
		logger.Info("Received a config", zap.Any("configs", configs))
		n.updateCache(&configs)
	})
}

func (n *Nats) Close() {
	n.conn.Close()
}

type NodeConfigNats struct {
	NodeType      string
	Id            string `bson:"_id"`
	WorkflowId    string
	Configuration json.RawMessage
	ClientId      string
}

type NodeConfig struct {
	NodeType      string
	ID            string `bson:"_id"`
	WorkflowID    string
	Configuration proto.Message
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

func (s *NodeConfigNats) decode(config proto.Message) (*NodeConfig, error) {
	//newRef := reflect.New(reflect.TypeOf(configRef))
	//newConf := newRef.Interface()

	//err := protojson.Unmarshal(s.configuration, &newConf)
	err := protojson.Unmarshal(s.Configuration, config)
	if err != nil {
		return nil, err
	}

	return &NodeConfig{
		NodeType:      s.NodeType,
		ID:            s.Id,
		WorkflowID:    s.WorkflowId,
		Configuration: config,
		ClientID:      s.ClientId,
	}, nil
}

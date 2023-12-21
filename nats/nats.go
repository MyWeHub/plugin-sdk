package nats

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MyWeHub/plugin-sdk/util"
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

type ListenerOptions struct {
	updateFunc func(nc *NodeConfig)
	removeFunc func(nc *NodeConfig)
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

func (n *Nats) Listen(ctx context.Context, opts ...*ListenerOptions) { // TODO: in previous files, Request() used EncodedConn but Subscribe() used natsConn [* not encoded] !
	ctx, span := tracer.Start(ctx, "Request plugin configuration")
	defer span.End()

	updateFunc := func(nc *NodeConfig) {}
	removeFunc := func(nc *NodeConfig) {}

	if opts != nil && len(opts) != 0 {
		for _, opt := range opts {
			if opt.updateFunc != nil {
				updateFunc = opt.updateFunc
			}
			if opt.removeFunc != nil {
				removeFunc = opt.removeFunc
			}
		}
	}

	pluginName := util.GetEnv("PLUGIN_NAME", false, "", true)

	req := fmt.Sprintf("requestPluginConfig.%s", pluginName)
	span.SetAttributes(attribute.String("topic", req))

	natsConfigs := make([]NodeConfigNats, 0)
	if err := n.conn.Request(req, "", &natsConfigs, 10*time.Second); err != nil {
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

	n.updateCache(&configs, updateFunc)

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
		n.updateCache(&configs, updateFunc)
	})

	n.conn.Subscribe(fmt.Sprintf("unpublishConfiguration.%s.*", pluginName), func(m *nats.Msg) {
		ctx, span = tracer.Start(ctx, "Received unpublished event")
		defer span.End()

		if err := json.Unmarshal(m.Data, &natsConfigs); err != nil {
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

		n.removeFromCache(&configs, removeFunc)
		logger.Info("Received unpublished event", zap.Any("configs", configs))
	})
}

func (n *Nats) updateCache(configs *[]NodeConfig, customFunc func(nc *NodeConfig)) {
	for _, nodeConfig := range *configs {
		nc := nodeConfig
		n.Cache[nodeConfig.ID] = &nc

		if customFunc != nil {
			customFunc(&nc)
		}
	}
}

func (n *Nats) removeFromCache(configs *[]NodeConfig, customFunc func(nc *NodeConfig)) {
	for _, nodeConfig := range *configs {
		if _, found := n.Cache[nodeConfig.ID]; found {
			if customFunc != nil {
				customFunc(&nodeConfig)
			}

			delete(n.Cache, nodeConfig.ID)
		}
	}
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

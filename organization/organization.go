package organization

import (
	"context"
	"errors"
	pbsc "github.com/MyWeHub/plugin-sdk/gen/serviceConnection"
	grpcOtel "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var (
	logger *zap.Logger
	tracer trace.Tracer
)

const (
	externalURL = "grpc-uat.weconnecthub.net:80"
	internalURL = "organization-management-server:6852"
	moduleName  = "organization-management-server-api"
)

type Organization struct {
	conn *grpc.ClientConn
	pbsc.ConnectionServiceClient
}

type Options struct {
	ExternalRequest bool
}

func SetTelemetry(l *zap.Logger, t trace.Tracer) {
	logger = l
	tracer = t
}

func New(ctx context.Context, opts ...*Options) (*Organization, context.Context, error) {
	url := internalURL
	if opts != nil && len(opts) > 0 {
		switch {
		case opts[0].ExternalRequest:
			url = externalURL
		}
	}

	conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithStatsHandler(grpcOtel.NewClientHandler()))
	if err != nil {
		logger.Error("Dial service-connection", zap.Error(err), zap.String("url", url))
		return nil, nil, err
	}

	orgCtx, err := configureCtx(ctx)
	if err != nil {
		return nil, nil, err
	}

	client := pbsc.NewConnectionServiceClient(conn)

	return &Organization{
		ConnectionServiceClient: client,
		conn:                    conn,
	}, orgCtx, nil
}

func (c *Organization) Close() error {
	return c.conn.Close()
}

func configureCtx(ctx context.Context) (context.Context, error) {
	jwt, ok := ctx.Value("token").(string)
	if !ok {
		logger.Error("organization-client: token not found in context")
		return nil, errors.New("organization-client: token not found in context")
	}

	md := metadata.New(map[string]string{
		"module":        moduleName,
		"authorization": "bearer " + jwt,
	})

	return metadata.NewOutgoingContext(ctx, md), nil
}

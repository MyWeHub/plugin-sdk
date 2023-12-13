package connectionService

import (
	"context"
	pbsc "dev.azure.com/WeConnectTechnology/ExchangeHub/_git/wehublib.git/gen/serviceConnection"
	"errors"
	"fmt"
	grpcOtel "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	logger *zap.Logger
	tracer trace.Tracer
)

const (
	externalURL = "grpc-uat.weconnecthub.net:80"
	internalURL = "service-connection:6852"
	moduleName  = "service-connection-api"
)

type ConnectionService struct {
	conn   *grpc.ClientConn
	client pbsc.ConnectionServiceClient
}

type Options struct {
	ExternalRequest bool
}

type connectionMessage struct {
	message *pbsc.ConnectionsMessage
}

func SetTelemetry(l *zap.Logger, t trace.Tracer) {
	logger = l
	tracer = t
}

func New(ctx context.Context, opts ...*Options) (*ConnectionService, error) {
	url := internalURL
	if opts != nil && len(opts) > 0 {
		switch {
		case opts[0].ExternalRequest:
			url = externalURL
		}
	}

	conn, err := grpc.DialContext(ctx, url, grpc.WithInsecure(), grpc.WithUnaryInterceptor(grpcOtel.UnaryClientInterceptor()))
	if err != nil {
		logger.Error("Dial service-connection", zap.Error(err), zap.String("url", url))
		return nil, err
	}

	client := pbsc.NewConnectionServiceClient(conn)

	return &ConnectionService{
		conn:   conn,
		client: client,
	}, nil
}

func (cs *ConnectionService) Close() error {
	return cs.conn.Close()
}

func (cs *ConnectionService) GetConnection(ctx context.Context, id string) (*connectionMessage, error) {
	if cs.client == nil {
		return nil, errors.New("service-connection client connection is empty")
	}

	jwt, ok := ctx.Value("token").(string)
	if !ok {
		logger.Error("service-connection: token not found in context")
		return nil, errors.New("service-connection: token not found in context")
	}

	md := metadata.New(map[string]string{
		"module":        moduleName,
		"authorization": "bearer " + jwt,
		//"authorization": "bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6Ilg1ZVhrNHh5b2pORnVtMWtsMll0djhkbE5QNC1jNTdkTzZRR1RWQndhTmsifQ.eyJleHAiOjE2NTYwNjUwODcsIm5iZiI6MTY1NjA2MTQ4NywidmVyIjoiMS4wIiwiaXNzIjoiaHR0cHM6Ly93ZWNvbm5lY3RodWIuYjJjbG9naW4uY29tLzZjOTdhZDE1LThkODYtNDg2My1hMWI0LTZhODU1ODE1MDUyZC92Mi4wLyIsInN1YiI6ImE5MjUwZTQ5LTgxMjQtNDViZC1hOTZkLWRhMzM5OTkxMmU2NCIsImF1ZCI6Ijc5MzkzMmUyLTlkNjctNDIwYy04NjFiLWEyNjZjYTc2YmIzYSIsImlhdCI6MTY1NjA2MTQ4NywiYXV0aF90aW1lIjoxNjU2MDYxNDg2LCJvaWQiOiJhOTI1MGU0OS04MTI0LTQ1YmQtYTk2ZC1kYTMzOTk5MTJlNjQiLCJuYW1lIjoiUGF0Y2h3b3JrIFRlc3RjbGllbnQiLCJnaXZlbl9uYW1lIjoiUGF0Y2h3b3JrIiwiZmFtaWx5X25hbWUiOiJUZXN0Y2xpZW50IiwiZXh0ZW5zaW9uX2NsaWVudGlkIjoiNjI1NTFjMDFkNDVlOGE0NDk0MDY3M2YyIiwiZXh0ZW5zaW9uX3dvcmthdG9jbGllbnRpZCI6IjYxMTM2MSIsImV4dGVuc2lvbl9pc0FkbWluIjp0cnVlLCJlbWFpbHMiOlsicGF0Y2h3b3JrY2xpZW50QHdlY29ubmVjdGh1Yi5jb20iXSwidGZwIjoiQjJDXzFfU2lnblVwU2lnbkluIiwiYXRfaGFzaCI6IjVNZmpMUWxLZnNndEdSYkllX2xFY0EifQ.gdMzwcftq6IsBiMAutwidPX2T2FyGGH2SPYYbhIABRdJFq87N7V8x15t-_almK9kL2K0E9yZDikgtfsaxIPaxwWBh-djk3LrBWvdf54bJMVb0PRa6pzfHmsyb2R9EHVpb6ty1-IzgY7DpE7YC7wbu2YqnswMT4UomygE6adN89bo_O4DJFImItvErnWP4jLOSyplhhb4zlE3OuSV5VV34UMpMzYJhhnTE3E0-bl_9zsNsGtLteo7CjB0cMf1W8NiRotKkZwhxq8uEXQxFVuthl-qPWDq70yFBmgQv5ZKJbydP4tkbEHQNtXls9zViuqXiSe54YCm8yVxczjc5EQKQA",
	})

	ctx = metadata.NewOutgoingContext(ctx, md)

	res, err := cs.client.Get(ctx, &pbsc.IdMessage{Id: id})
	if err != nil {
		return nil, err
	}

	return &connectionMessage{message: res}, nil
}

func (cm *connectionMessage) ToSFTP() (*pbsc.SFTPConnection, error) {
	conn, ok := cm.message.Connection.ConnectionOneof.(*pbsc.Connection_SftpConnection)
	if !ok {
		logger.Error("service-connection: can't convert ConnectionOneOf", zap.Any("type", pbsc.ConnectionType_CONNECTION_SFTP))
		return nil, fmt.Errorf("service-connection: can't convert ConnectionOneOf to type '%v'", pbsc.ConnectionType_CONNECTION_SFTP)
	}

	return conn.SftpConnection, nil
}

func (cm *connectionMessage) ToMongo() (*pbsc.MongoConnection, error) {
	conn, ok := cm.message.Connection.ConnectionOneof.(*pbsc.Connection_MongoConnection)
	if !ok {
		logger.Error("service-connection: can't convert ConnectionOneOf", zap.Any("type", pbsc.ConnectionType_CONNECTION_MONGO))
		return nil, fmt.Errorf("service-connection: can't convert ConnectionOneOf to type '%v'", pbsc.ConnectionType_CONNECTION_MONGO)
	}

	return conn.MongoConnection, nil
}

func (cm *connectionMessage) ToAMQP() (*pbsc.AMQPConnection, error) {
	conn, ok := cm.message.Connection.ConnectionOneof.(*pbsc.Connection_AmqpConnection)
	if !ok {
		logger.Error("service-connection: can't convert ConnectionOneOf", zap.Any("type", pbsc.ConnectionType_CONNECTION_AMQP))
		return nil, fmt.Errorf("service-connection: can't convert ConnectionOneOf to type '%v'", pbsc.ConnectionType_CONNECTION_AMQP)
	}

	return conn.AmqpConnection, nil
}

func (cm *connectionMessage) ToKafka() (*pbsc.KafkaConnection, error) {
	conn, ok := cm.message.Connection.ConnectionOneof.(*pbsc.Connection_KafkaConnection)
	if !ok {
		logger.Error("service-connection: can't convert ConnectionOneOf", zap.Any("type", pbsc.ConnectionType_CONNECTION_KAFKA))
		return nil, fmt.Errorf("service-connection: can't convert ConnectionOneOf to type '%v'", pbsc.ConnectionType_CONNECTION_KAFKA)
	}

	return conn.KafkaConnection, nil
}

func (cm *connectionMessage) ToHTTP() (*pbsc.HTTPConnection, error) {
	conn, ok := cm.message.Connection.ConnectionOneof.(*pbsc.Connection_HttpConnection)
	if !ok {
		logger.Error("service-connection: can't convert ConnectionOneOf", zap.Any("type", pbsc.ConnectionType_CONNECTION_HTTP))
		return nil, fmt.Errorf("service-connection: can't convert ConnectionOneOf to type '%v'", pbsc.ConnectionType_CONNECTION_HTTP)
	}

	return conn.HttpConnection, nil
}

func (cm *connectionMessage) ToRedis() (*pbsc.RedisConnection, error) {
	conn, ok := cm.message.Connection.ConnectionOneof.(*pbsc.Connection_RedisConnection)
	if !ok {
		logger.Error("service-connection: can't convert ConnectionOneOf", zap.Any("type", pbsc.ConnectionType_CONNECTION_REDIS))
		return nil, fmt.Errorf("service-connection: can't convert ConnectionOneOf to type '%v'", pbsc.ConnectionType_CONNECTION_REDIS)
	}

	return conn.RedisConnection, nil
}

func (cm *connectionMessage) ToTwilio() (*pbsc.TwilioConnection, error) {
	conn, ok := cm.message.Connection.ConnectionOneof.(*pbsc.Connection_TwilioConnection)
	if !ok {
		logger.Error("service-connection: can't convert ConnectionOneOf", zap.Any("type", pbsc.ConnectionType_CONNECTION_TWILIO))
		return nil, fmt.Errorf("service-connection: can't convert ConnectionOneOf to type '%v'", pbsc.ConnectionType_CONNECTION_TWILIO)
	}

	return conn.TwilioConnection, nil
}

func (cm *connectionMessage) ToSendgrid() (*pbsc.SendGridConnection, error) {
	conn, ok := cm.message.Connection.ConnectionOneof.(*pbsc.Connection_SendgridConnection)
	if !ok {
		logger.Error("service-connection: can't convert ConnectionOneOf", zap.Any("type", pbsc.ConnectionType_CONNECTION_SENDGRID))
		return nil, fmt.Errorf("service-connection: can't convert ConnectionOneOf to type '%v'", pbsc.ConnectionType_CONNECTION_SENDGRID)
	}

	return conn.SendgridConnection, nil
}

func (cm *connectionMessage) ToCosmosDB() (*pbsc.CosmosDBConnection, error) {
	conn, ok := cm.message.Connection.ConnectionOneof.(*pbsc.Connection_CosmosdbConnection)
	if !ok {
		logger.Error("service-connection: can't convert ConnectionOneOf", zap.Any("type", pbsc.ConnectionType_CONNECTION_COSMOSDB))
		return nil, fmt.Errorf("service-connection: can't convert ConnectionOneOf to type '%v'", pbsc.ConnectionType_CONNECTION_COSMOSDB)
	}

	return conn.CosmosdbConnection, nil
}

func (cm *connectionMessage) ToMySQL() (*pbsc.MySQLConnection, error) {
	conn, ok := cm.message.Connection.ConnectionOneof.(*pbsc.Connection_MysqlConnection)
	if !ok {
		logger.Error("service-connection: can't convert ConnectionOneOf", zap.Any("type", pbsc.ConnectionType_CONNECTION_MYSQL))
		return nil, fmt.Errorf("service-connection: can't convert ConnectionOneOf to type '%v'", pbsc.ConnectionType_CONNECTION_MYSQL)
	}

	return conn.MysqlConnection, nil
}

func (cm *connectionMessage) ToMSSQL() (*pbsc.MsSQLConnection, error) {
	conn, ok := cm.message.Connection.ConnectionOneof.(*pbsc.Connection_MssqlConnection)
	if !ok {
		logger.Error("service-connection: can't convert ConnectionOneOf", zap.Any("type", pbsc.ConnectionType_CONNECTION_MSSQL))
		return nil, fmt.Errorf("service-connection: can't convert ConnectionOneOf to type '%v'", pbsc.ConnectionType_CONNECTION_MSSQL)
	}

	return conn.MssqlConnection, nil
}

func (cm *connectionMessage) ToPostgres() (*pbsc.PostgresConnection, error) {
	conn, ok := cm.message.Connection.ConnectionOneof.(*pbsc.Connection_PostgresConnection)
	if !ok {
		logger.Error("service-connection: can't convert ConnectionOneOf", zap.Any("type", pbsc.ConnectionType_CONNECTION_POSTGRES))
		return nil, fmt.Errorf("service-connection: can't convert ConnectionOneOf to type '%v'", pbsc.ConnectionType_CONNECTION_POSTGRES)
	}

	return conn.PostgresConnection, nil
}

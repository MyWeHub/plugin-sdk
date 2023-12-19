package testing

import (
	"context"
	wehublib "github.com/MyWeHub/plugin-sdk"
	pb "github.com/MyWeHub/plugin-sdk/gen/pluginrunner"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
	"net"
)

var (
	logger *zap.Logger
	tracer trace.Tracer
)

const bufSize = 1024 * 1024

var lis = bufconn.Listen(bufSize)

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

type testing struct {
	conn   *grpc.ClientConn
	client pb.PluginRunnerServiceClient
}

type IService interface {
	Process(ctx context.Context, in *structpb.Struct, conf proto.Message, action int32, workflowData string) (*structpb.Struct, error)
}

func New(ctx context.Context, is IService) *testing {
	t := wehublib.NewTelemetry()
	defer t.ShutdownTracer(ctx)
	defer t.SyncLogger()

	logger = t.GetLogger()
	tracer = t.GetTracer()

	server := wehublib.NewServer()
	server.SetNewGRPC()
	server.RegisterServer(nil, is)

	go func() {
		if err := server.ServeTest(lis); err != nil {
			panic(err)
		}
	}()

	return &testing{}
}

func (t *testing) NewClient(ctx context.Context) pb.PluginRunnerServiceClient {
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		logger.Error("Failed to dial bufnet: ", zap.Error(err))
	}

	client := pb.NewPluginRunnerServiceClient(conn)

	t.conn = conn
	t.client = client

	return client
}

func (t *testing) Close() error {
	return t.conn.Close()
}

var token = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6Ilg1ZVhrNHh5b2pORnVtMWtsMll0djhkbE5QNC1jNTdkTzZRR1RWQndhTmsifQ.eyJleHAiOjE2NTYwNjUwODcsIm5iZiI6MTY1NjA2MTQ4NywidmVyIjoiMS4wIiwiaXNzIjoiaHR0cHM6Ly93ZWNvbm5lY3RodWIuYjJjbG9naW4uY29tLzZjOTdhZDE1LThkODYtNDg2My1hMWI0LTZhODU1ODE1MDUyZC92Mi4wLyIsInN1YiI6ImE5MjUwZTQ5LTgxMjQtNDViZC1hOTZkLWRhMzM5OTkxMmU2NCIsImF1ZCI6Ijc5MzkzMmUyLTlkNjctNDIwYy04NjFiLWEyNjZjYTc2YmIzYSIsImlhdCI6MTY1NjA2MTQ4NywiYXV0aF90aW1lIjoxNjU2MDYxNDg2LCJvaWQiOiJhOTI1MGU0OS04MTI0LTQ1YmQtYTk2ZC1kYTMzOTk5MTJlNjQiLCJuYW1lIjoiUGF0Y2h3b3JrIFRlc3RjbGllbnQiLCJnaXZlbl9uYW1lIjoiUGF0Y2h3b3JrIiwiZmFtaWx5X25hbWUiOiJUZXN0Y2xpZW50IiwiZXh0ZW5zaW9uX2NsaWVudGlkIjoiNjI1NTFjMDFkNDVlOGE0NDk0MDY3M2YyIiwiZXh0ZW5zaW9uX3dvcmthdG9jbGllbnRpZCI6IjYxMTM2MSIsImV4dGVuc2lvbl9pc0FkbWluIjp0cnVlLCJlbWFpbHMiOlsicGF0Y2h3b3JrY2xpZW50QHdlY29ubmVjdGh1Yi5jb20iXSwidGZwIjoiQjJDXzFfU2lnblVwU2lnbkluIiwiYXRfaGFzaCI6IjVNZmpMUWxLZnNndEdSYkllX2xFY0EifQ.gdMzwcftq6IsBiMAutwidPX2T2FyGGH2SPYYbhIABRdJFq87N7V8x15t-_almK9kL2K0E9yZDikgtfsaxIPaxwWBh-djk3LrBWvdf54bJMVb0PRa6pzfHmsyb2R9EHVpb6ty1-IzgY7DpE7YC7wbu2YqnswMT4UomygE6adN89bo_O4DJFImItvErnWP4jLOSyplhhb4zlE3OuSV5VV34UMpMzYJhhnTE3E0-bl_9zsNsGtLteo7CjB0cMf1W8NiRotKkZwhxq8uEXQxFVuthl-qPWDq70yFBmgQv5ZKJbydP4tkbEHQNtXls9zViuqXiSe54YCm8yVxczjc5EQKQA"

func AppendIncomingTestToken(ctx context.Context) context.Context {
	md := metadata.New(map[string]string{
		"authorization": "bearer " + token,
	})

	return metadata.NewOutgoingContext(ctx, md)
}

func AppendInterceptorTestToken(ctx context.Context) context.Context {
	return context.WithValue(ctx, "token", token)
}

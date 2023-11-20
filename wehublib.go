package wehublib

import (
	"context"
	pb "dev.azure.com/WeConnectTechnology/ExchangeHub/_git/wehublib.git/gen/pluginrunner"
	"dev.azure.com/WeConnectTechnology/ExchangeHub/_git/wehublib.git/telemetry"
	"dev.azure.com/WeConnectTechnology/ExchangeHub/_git/wehublib.git/util"
	"errors"
	"fmt"
	"github.com/amsokol/mongo-go-driver-protobuf/pmongo"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcAuth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcTags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	grpcOtel "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

// TODO: maybe add options to control which interceptor to add or not
// TODO: submodule?
// TODO: write unit tests!
// TODO: add README.md and documentation for godoc
// TODO: grpc status pkg!
// TODO:

type server struct {
	server       *grpc.Server
	logger       *zap.Logger
	grpcPort     string
	httpPort     string
	jwtAuthFunc  func(ctx context.Context) (context.Context, error)
	recoveryFunc func(interface{}) error
}

type ServerOptions struct {
	DisableHTTP      bool
	GracefulShutdown bool
}

type GRPCOptions struct {
	TagsInterceptor       bool
	OtelInterceptor       bool
	PrometheusInterceptor bool
	ZapLoggerInterceptor  bool
	AuthInterceptor       bool
	RecoveryInterceptor   bool
	MaxReceiveSize        int
	MaxSendSize           int
}

func NewServer(t *telemetry.Telemetry) *server {
	return &server{
		logger:       t.GetLogger(),
		grpcPort:     util.GetEnv("GRPC_PORT", true, "6852", false),
		httpPort:     util.GetEnv("HTTP_PORT", true, "3000", false),
		jwtAuthFunc:  defaultAuthFunc,
		recoveryFunc: defaultRecoveryFunc,
	}
}

func (s *server) SetServiceServer(srv pb.PluginRunnerServiceServer) {
	if s.server == nil {
		panic(errors.New("grpc server must be initialized before setting service server"))
	}
	pb.RegisterPluginRunnerServiceServer(s.server, srv)
}

func (s *server) SetCustomGRPCPort(p string) {
	s.grpcPort = p
}

func (s *server) SetCustomHTTPPort(p string) {
	s.httpPort = p
}

func (s *server) SetCustomJwtHandler(handler func(ctx context.Context) (context.Context, error)) {
	s.jwtAuthFunc = handler
}

func (s *server) SetCustomRecoveryHandler(handler func(interface{}) error) {
	s.recoveryFunc = handler
}

func (s *server) SetNewGRPC() *server {
	if s.logger == nil {
		panic(errors.New("logger not set. please use 'SetLogger' method before initializing server"))
	}

	recoveryOpts := []grpcRecovery.Option{
		grpcRecovery.WithRecoveryHandler(s.recoveryFunc),
	}

	streamInterceptor := grpc.StreamInterceptor(
		grpcMiddleware.ChainStreamServer(
			grpcTags.StreamServerInterceptor(),
			grpcOtel.StreamServerInterceptor(),
			grpcPrometheus.StreamServerInterceptor,
			grpcZap.StreamServerInterceptor(s.logger),
			grpcAuth.StreamServerInterceptor(s.jwtAuthFunc),
			grpcRecovery.StreamServerInterceptor(),
			grpcRecovery.StreamServerInterceptor(recoveryOpts...)))
	unaryInterceptor := grpc.UnaryInterceptor(
		grpcMiddleware.ChainUnaryServer(
			grpcTags.UnaryServerInterceptor(),
			grpcOtel.UnaryServerInterceptor(),
			grpcPrometheus.UnaryServerInterceptor,
			grpcZap.UnaryServerInterceptor(s.logger),
			grpcAuth.UnaryServerInterceptor(s.jwtAuthFunc),
			grpcRecovery.UnaryServerInterceptor(),
			grpcRecovery.UnaryServerInterceptor(recoveryOpts...)))

	s.server = grpc.NewServer(streamInterceptor, unaryInterceptor)
	return s
}

func (s *server) SetNewCustomGRPC(opts ...*GRPCOptions) *server {
	if len(opts) == 0 || opts == nil {
		return s.SetNewGRPC()
	}

	stream := make([]grpc.StreamServerInterceptor, 0)
	unary := make([]grpc.UnaryServerInterceptor, 0)

	x := opts[0]

	if x.TagsInterceptor {
		unary = append(unary, grpcTags.UnaryServerInterceptor())
		stream = append(stream, grpcTags.StreamServerInterceptor())
	}
	if x.OtelInterceptor {
		unary = append(unary, grpcOtel.UnaryServerInterceptor())
		stream = append(stream, grpcOtel.StreamServerInterceptor())
	}
	if x.PrometheusInterceptor {
		unary = append(unary, grpcPrometheus.UnaryServerInterceptor)
		stream = append(stream, grpcPrometheus.StreamServerInterceptor)
	}
	if x.ZapLoggerInterceptor {
		if s.logger == nil {
			panic(errors.New("logger not set. please use 'SetLogger' method before initializing server"))
		}

		unary = append(unary, grpcZap.UnaryServerInterceptor(s.logger))
		stream = append(stream, grpcZap.StreamServerInterceptor(s.logger))
	}
	if x.AuthInterceptor {
		unary = append(unary, grpcAuth.UnaryServerInterceptor(s.jwtAuthFunc))
		stream = append(stream, grpcAuth.StreamServerInterceptor(s.jwtAuthFunc))
	}
	if x.RecoveryInterceptor {
		recoveryOpts := []grpcRecovery.Option{
			grpcRecovery.WithRecoveryHandler(s.recoveryFunc),
		}

		unary = append(unary, grpcRecovery.UnaryServerInterceptor(recoveryOpts...))
		stream = append(stream, grpcRecovery.StreamServerInterceptor(recoveryOpts...))
	}

	fmt.Println("------------")
	fmt.Println(len(unary))
	fmt.Println(len(stream))
	fmt.Println("------------")

	serverOpts := make([]grpc.ServerOption, 0)
	serverOpts = append(serverOpts, grpc.StreamInterceptor(grpcMiddleware.ChainStreamServer(stream...)))
	serverOpts = append(serverOpts, grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(unary...)))

	if x.MaxSendSize > 0 {
		serverOpts = append(serverOpts, grpc.MaxSendMsgSize(x.MaxSendSize))
	}
	if x.MaxReceiveSize > 0 {
		serverOpts = append(serverOpts, grpc.MaxRecvMsgSize(x.MaxReceiveSize))
	}

	fmt.Println("------------")
	fmt.Println(len(serverOpts))
	fmt.Println("------------")

	s.server = grpc.NewServer(serverOpts...)

	return s
}

func (s *server) Serve(opts ...*ServerOptions) {
	if s.server == nil {
		panic(errors.New("server not initialized"))
	}

	flagServeHttp := true
	flagGracefulShutdown := true

	if len(opts) > 0 && opts[0] != nil {
		switch {
		case opts[0].DisableHTTP:
			flagServeHttp = false
		case !opts[0].GracefulShutdown:
			flagGracefulShutdown = false
		}
	}

	if flagServeHttp {
		s.logger.Info("HTTP server started", zap.String("port", s.httpPort))
		s.serveHTTP()
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", s.grpcPort))
	if err != nil {
		panic(err)
	}

	if flagGracefulShutdown {

		idleConnsClosed := make(chan struct{})

		go func() {
			sigint := make(chan os.Signal, 1)

			// interrupt signal sent from terminal
			signal.Notify(sigint, os.Interrupt)
			// sigterm signal sent from kubernetes
			signal.Notify(sigint, syscall.SIGTERM)

			<-sigint

			// We received an interrupt signal, shut down.
			s.logger.Info("GRPC server gracefully shutdown")
			s.server.GracefulStop()
			close(idleConnsClosed)
		}()

		s.logger.Info("GRPC server started", zap.String("port", s.grpcPort))
		go func() {
			if err := s.server.Serve(lis); err != nil {
				panic(err)
			}
		}()

		<-idleConnsClosed

	} else {
		s.logger.Info("GRPC server started", zap.String("port", s.grpcPort))
		if err = s.server.Serve(lis); err != nil {
			panic(err)
		}
	}
}

func (s *server) ServeTest(lis net.Listener) error {
	return s.server.Serve(lis)
}

func (s *server) serveHTTP() {
	app := fiber.New()
	app.Static("/", "/go/bin/public", fiber.Static{Compress: true})
	go app.Listen(fmt.Sprintf(":%s", s.httpPort))
}

var (
	defaultRecoveryFunc = func(p interface{}) (err error) {
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}

	defaultAuthFunc = func(ctx context.Context) (context.Context, error) {
		token, err := grpcAuth.AuthFromMD(ctx, "bearer")
		if err != nil {
			return nil, err
		}

		claims := jwt.MapClaims{}
		parser := jwt.Parser{}

		tokenInfo, _, err := parser.ParseUnverified(token, claims)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
		}

		// WARNING: in production define your own type to avoid context collisions
		ctx = context.WithValue(ctx, "tokenInfo", tokenInfo)
		ctx = context.WithValue(ctx, "token", token)

		if clientID, ok := claims["extension_clientid"]; ok {
			oid, err := primitive.ObjectIDFromHex(clientID.(string))
			if err != nil {
				log.Println("convert string clientID to ObjectID:", err)
			}
			poid := pmongo.NewObjectId(oid)

			grpcTags.Extract(ctx).Set("auth.clientId", clientID)
			ctx = context.WithValue(ctx, "clientId", poid)
		} else {
			ctx = context.WithValue(ctx, "clientId", &pmongo.ObjectId{})
		}

		if isAdmin, ok := claims["extension_isAdmin"]; ok {
			grpcTags.Extract(ctx).Set("auth.superAdmin", isAdmin.(bool))
			ctx = context.WithValue(ctx, "superAdmin", isAdmin.(bool))
		} else {
			ctx = context.WithValue(ctx, "superAdmin", false)
		}

		grpcTags.Extract(ctx).Set("auth.sub", claims)

		return ctx, nil
	}
)

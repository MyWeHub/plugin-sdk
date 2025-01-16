package wehublib

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	pbEP "github.com/MyWeHub/plugin-sdk/gen/entrypointService"
	pb "github.com/MyWeHub/plugin-sdk/gen/pluginrunner"
	"github.com/MyWeHub/plugin-sdk/nats"
	"github.com/MyWeHub/plugin-sdk/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcAuth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcTags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	grpcOtel "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
)

// TODO: write unit tests!
// TODO: grpc status pkg!
// TODO: change package name from 'wehublib' to 'pluginSDK'
// TODO: update documentation
// TODO:

var (
	logger *zap.Logger
	tracer trace.Tracer
)

type server struct {
	HttpServer   *fiber.App
	GrpcServer   *grpc.Server
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

type IProcess interface {
	Process(ctx context.Context, in *structpb.Struct, conf proto.Message, action int32, workflowData string) (*pb.InputTestResponseV2, error)
}

type grpcServer struct {
	pb.UnimplementedPluginRunnerServiceServer
	nats       *nats.Nats
	service    IProcess
	configType proto.Message
}

func NewServer() *server {
	env := util.LoadEnvironment()
	switch env {
	case util.EnvPROD:
		log.Println("Initiating GrpcServer in 'PRODUCTION' mode")
	case util.EnvDEV:
		log.Println("Initiating GrpcServer in 'DEVELOPMENT' mode")
	}

	return &server{
		grpcPort:     util.GetEnv("GRPC_PORT", true, "6852", false),
		httpPort:     util.GetEnv("HTTP_PORT", true, "3000", false),
		jwtAuthFunc:  DefaultGrpcJwtAuthFunc,
		recoveryFunc: defaultRecoveryFunc,
	}
}

func (s *server) RegisterServer(n *nats.Nats, is IProcess, ct proto.Message) {
	if s.GrpcServer == nil {
		panic(errors.New("grpc server must be initialized before setting service server"))
	}

	pb.RegisterPluginRunnerServiceServer(s.GrpcServer, &grpcServer{nats: n, service: is, configType: ct})
}

func (s *server) RegisterEntrypointServer(srv pbEP.EntrypointServiceServer) {
	if s.GrpcServer == nil {
		panic(errors.New("grpc server must be initialized before setting service server"))
	}

	pbEP.RegisterEntrypointServiceServer(s.GrpcServer, srv)
}

func (s *server) GetGRPCServer() (*grpc.Server, error) {
	if s.GrpcServer == nil {
		return nil, errors.New("server is nil, please call this method after 'SetNewGRPC' method is called")
	}

	return s.GrpcServer, nil
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

func (s *server) SetNewGRPC(opts ...*GRPCOptions) *server {
	if len(opts) == 0 || opts == nil {
		return s.setDefaultGRPC()
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
		if logger == nil {
			panic(errors.New("logger not set. please use 'SetLogger' method before initializing server"))
		}

		unary = append(unary, grpcZap.UnaryServerInterceptor(logger))
		stream = append(stream, grpcZap.StreamServerInterceptor(logger))
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

	serverOpts := make([]grpc.ServerOption, 0)
	serverOpts = append(serverOpts, grpc.StreamInterceptor(grpcMiddleware.ChainStreamServer(stream...)))
	serverOpts = append(serverOpts, grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(unary...)))

	if x.MaxSendSize > 0 {
		serverOpts = append(serverOpts, grpc.MaxSendMsgSize(x.MaxSendSize))
	}
	if x.MaxReceiveSize > 0 {
		serverOpts = append(serverOpts, grpc.MaxRecvMsgSize(x.MaxReceiveSize))
	}

	s.GrpcServer = grpc.NewServer(serverOpts...)

	return s
}

func (s *server) setDefaultGRPC() *server {
	if logger == nil {
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
			grpcZap.StreamServerInterceptor(logger),
			grpcAuth.StreamServerInterceptor(s.jwtAuthFunc),
			grpcRecovery.StreamServerInterceptor(),
			grpcRecovery.StreamServerInterceptor(recoveryOpts...)))
	unaryInterceptor := grpc.UnaryInterceptor(
		grpcMiddleware.ChainUnaryServer(
			grpcTags.UnaryServerInterceptor(),
			grpcOtel.UnaryServerInterceptor(),
			grpcPrometheus.UnaryServerInterceptor,
			grpcZap.UnaryServerInterceptor(logger),
			grpcAuth.UnaryServerInterceptor(s.jwtAuthFunc),
			grpcRecovery.UnaryServerInterceptor(),
			grpcRecovery.UnaryServerInterceptor(recoveryOpts...)))

	s.GrpcServer = grpc.NewServer(streamInterceptor, unaryInterceptor)
	return s
}

func (s *server) Serve(opts ...*ServerOptions) {
	if s.GrpcServer == nil {
		panic(errors.New("server not initialized"))
	}

	if s.httpPort == s.grpcPort {
		panic(errors.New("http and grpc ports are the same"))
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
		logger.Info("HTTP server started", zap.String("port", s.httpPort))
		s.serveHTTP()
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", s.grpcPort))
	if err != nil {
		logger.Fatal("net.Listen", zap.Error(err))
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
			logger.Info("GRPC server gracefully shutdown")
			s.GrpcServer.GracefulStop()
			close(idleConnsClosed)
		}()

		logger.Info("GRPC server started", zap.String("port", s.grpcPort))
		go func() {
			if err := s.GrpcServer.Serve(lis); err != nil {
				panic(err)
			}
		}()

		<-idleConnsClosed

	} else {
		logger.Info("GRPC server started", zap.String("port", s.grpcPort))
		if err = s.GrpcServer.Serve(lis); err != nil {
			panic(err)
		}
	}
}

func (s *server) ServeTest(lis net.Listener) error {
	return s.GrpcServer.Serve(lis)
}

func (s *server) serveHTTP() {
	s.HttpServer = fiber.New()
	s.HttpServer.Use(compress.New())
	s.HttpServer.Static("/", "/go/bin/public", fiber.Static{Compress: true})
	http.Handle("/metrics", promhttp.Handler())
	s.HttpServer.Use(pprof.New())
	go http.ListenAndServe(":2112", nil)
	go s.HttpServer.Listen(fmt.Sprintf(":%s", s.httpPort))
}

var (
	defaultRecoveryFunc = func(p interface{}) (err error) {
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}

	DefaultGrpcJwtAuthFunc = func(ctx context.Context) (context.Context, error) {
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
			if clientID == "" {
				return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: ClientID is empty")
			}

			grpcTags.Extract(ctx).Set("auth.clientId", clientID)
			ctx = context.WithValue(ctx, "clientId", clientID)
		} else {
			return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: no ClientID found")
			//ctx = context.WithValue(ctx, "clientId", &pmongo.ObjectId{})
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

	DefaultRestJwtAuthMiddleware = func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing token",
			})
		}

		token = strings.TrimPrefix(token, "Bearer ")

		claims := jwt.MapClaims{}
		parser := jwt.Parser{}
		tkn, _, err := parser.ParseUnverified(token, claims)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": fmt.Errorf("Invalid token: " + err.Error()),
			})
		}

		if clientId, ok := claims["extension_clientid"]; !ok {
			if superAdmin, ok := claims["extension_superAdmin"]; !ok {
				if superAdmin.(bool) == true {
					c.Locals("auth.superAdmin", superAdmin)
					c.Locals("superAdmin", superAdmin)
				}
			} else {
				c.Locals("superAdmin", false)
			}
		} else {
			c.Locals("auth.clientId", clientId)
			c.Locals("clientId", clientId)
			c.Locals("superAdmin", false)
		}

		c.Locals("auth.sub", claims)
		c.Locals("tokenInfo", tkn)

		return c.Next()
	}
)

package main2

import (
	"context"
	"dev.azure.com/WeConnectTechnology/ExchangeHub/_git/wehublib.git"
	"dev.azure.com/WeConnectTechnology/ExchangeHub/_git/wehublib.git/nats"
	"dev.azure.com/WeConnectTechnology/ExchangeHub/_git/wehublib.git/telemetry"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/structpb"
)

var logger *zap.Logger

type service struct{}

func newService() *service {
	return &service{}
}

func (s *service) Process(ctx context.Context, in *structpb.Struct, conf interface{}, action int32, workflowData string) (*structpb.Struct, error) {
	return nil, nil
}

func main() {
	ctx := context.Background()

	// telemetry
	t := telemetry.New()
	defer t.ShutdownTracer(ctx)
	defer t.SyncLogger()

	//nats
	// TODO: send pointer to pbconf.config in NewNats, then save the decoded conf in the cache, and just type assert it when you get it
	n := nats.New(t)
	defer n.Close()
	n.Listen(ctx)
	if node, ok := n.Cache["input.NodeId"]; ok {
		fmt.Println(node.NodeType)
		fmt.Println(node.ID)
		fmt.Println(node.WorkflowID)
		//fmt.Println(node.DecodeConfig(nil))
	}

	// custom plugin handler

	// server
	server := wehublib.NewServer(t)
	server.SetCustomGRPCPort("6666")
	server.SetCustomHTTPPort("6666")
	//server.SetCustomJwtHandler()
	//server.SetCustomRecoveryHandler()
	server.SetNewGRPC(&wehublib.GRPCOptions{
		TagsInterceptor:       true,
		OtelInterceptor:       true,
		PrometheusInterceptor: true,
		ZapLoggerInterceptor:  true,
		AuthInterceptor:       true,
		RecoveryInterceptor:   true,
		MaxReceiveSize:        1024 * 1024 * 1024 * 1024,
		MaxSendSize:           1024 * 1024 * 1024 * 1024,
	})
	server.RegisterServer(ctx, newService())
	server.Serve(&wehublib.ServerOptions{DisableHTTP: false, GracefulShutdown: true})

	//
	//// connection service
	//ncs, err := cs.NewConnectionService(ctx, t, &cs.Options{ExternalRequest: true})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer ncs.Close()
	//
	//connection, err := ncs.GetConnection(testingLib.AppendInterceptorTestToken(ctx), "63464297-8f51-4094-96be-de25f9b44183")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//mongoConnection, err := connection.ToMongo()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("mongoConnection")
	//fmt.Println(mongoConnection)
}

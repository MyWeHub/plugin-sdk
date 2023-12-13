package main2

import (
	"context"
	"dev.azure.com/WeConnectTechnology/ExchangeHub/_git/wehublib.git"
	confPB "dev.azure.com/WeConnectTechnology/ExchangeHub/_git/wehublib.git/gen/configuration"
	"dev.azure.com/WeConnectTechnology/ExchangeHub/_git/wehublib.git/nats"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
)

var logger *zap.Logger

type service struct{}

func newService() *service {
	return &service{}
}

func (s *service) Process(ctx context.Context, in *structpb.Struct, conf proto.Message, action int32, workflowData string) (*structpb.Struct, error) {
	return nil, nil
}

func main() {
	ctx := context.Background()

	// telemetry
	t := wehublib.NewTelemetry()
	defer t.ShutdownTracer(ctx)
	defer t.SyncLogger()

	//nats
	n := nats.New(&confPB.Configuration{})
	defer n.Close()
	n.Listen(ctx)
	if node, ok := n.Cache["input.NodeId"]; ok {
		fmt.Println(node.NodeType)
		fmt.Println(node.ID)
		fmt.Println(node.WorkflowID)
		//fmt.Println(node.DecodeConfig(nil))
	}

	// connection service
	/*ncs, err := cs.New(ctx, &cs.Options{ExternalRequest: true})
	if err != nil {
		log.Fatal(err)
	}
	defer ncs.Close()

	connection, err := ncs.GetConnection(testingLib.AppendInterceptorTestToken(ctx), "63464297-8f51-4094-96be-de25f9b44183")
	if err != nil {
		log.Fatal(err)
	}

	mongoConnection, err := connection.ToMongo()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongoConnection")
	fmt.Println(mongoConnection)*/

	// server
	server := wehublib.NewServer()
	server.SetCustomGRPCPort("6666")
	server.SetCustomHTTPPort("6667")
	//server.SetCustomJwtHandler()
	//server.SetCustomRecoveryHandler()
	server.SetNewGRPC()
	server.RegisterServer(n, newService())
	server.Serve()
}

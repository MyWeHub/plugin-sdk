package main

import (
	"context"
	"fmt"
	pb "wehublib/gen/pluginrunner"
	"wehublib/nats"
	"wehublib/telemetry"
)

type serviceServer struct {
	pb.UnimplementedPluginRunnerServiceServer
}

func (c *serviceServer) RunTest(ctx context.Context, input *pb.InputTestRequest) (*pb.InputTestResponse, error) {
	fmt.Println("------------------------SHIT!")
	return &pb.InputTestResponse{}, nil
}

func main() {
	ctx := context.Background()

	// telemetry
	t := telemetry.NewTelemetry()
	defer t.ShutdownTracer(ctx)
	defer t.SyncLogger()

	// server
	/*server := wehublib.NewServer(t)
	server.SetNewGRPC()
	server.SetServiceServer(&serviceServer{})
	server.Serve(&wehublib.ServerOptions{DisableHTTP: false, GracefulShutdown: true})*/

	// nats
	n := nats.NewNats(t)
	defer n.Close()
	n.Listen(ctx)
	if node, ok := n.Cache["input.NodeId"]; ok {
		fmt.Println(node.NodeType)
		fmt.Println(node.ID)
		fmt.Println(node.WorkflowID)
		//fmt.Println(node.DecodeConfig(nil))
	}

	// connection service
	/*ncs, err := cs.NewConnectionService(ctx, t, &cs.Options{ExternalRequest: true})
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
}

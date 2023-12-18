# Wehublib
This is the repository of all the boilerplate and useful pieces of code that we use in WeHub Fusion Company. This repository is meant to be used in every plugin and also increase the speed of plugin development.

### Note: 
Using this library obligates you to use it's own generated protocol buffer files, instead of generating it separately for every plugin.

## Get Package
```shell
go get github.com/MyWeHub/plugin-sdk
```

## Telemetry
Every package in the library needs the telemetry instance, so we need to create one before we do anything else:
```go
package main

import "github.com/MyWeHub/plugin-sdk/telemetry"

func main() {
	t := telemetry.NewTelemetry()
	defer t.ShutdownTracer(ctx)
	defer t.SyncLogger()
}
```

## Server
In order to initialize the GRPC and HTTP server that every plugin needs, you can simply use the code below:
```go
package main

import "github.com/MyWeHub/plugin-sdk"

func main() {
	server := wehublib.NewServer(t)
	server.SetNewGRPC().SetServiceServer(&serviceServer{})
	server.SetCustomGRPCPort("6666")
	server.SetCustomHTTPPort("6666")
	server.SetCustomJwtHandler(...)
	server.SetCustomRecoveryHandler(...)
	server.Serve(&wehublib.ServerOptions{DisableHTTP: false, GracefulShutdown: true})
}
```
## Connection Service
Invoke Connection Service as described below:
```go
package main

import cs "github.com/MyWeHub/plugin-sdk/connectionService"

func main() {
	ncs, err := cs.NewConnectionService(ctx, t, &cs.Options{ExternalRequest: true})
	if err != nil {
		log.Fatal(err)
	}
	defer ncs.Close()

	// ctx must contain "token" metadata
	connection, err := ncs.GetConnection(ctx, "id")
	if err != nil {
		log.Fatal(err)
	}

	mongoConnection, err := connection.ToMongo()
	if err != nil {
		log.Fatal(err)
	}
}
```

## Nats
Every plugin uses Nats message broker in order to receive its configuration when in a published workflow. All useful codes can be used as below:

```go
package main

import (
	"github.com/MyWeHub/plugin-sdk/nats"
	"log"
)

func main() {
	n := nats.NewNats(t)
	defer n.Close()
	n.Listen(ctx)
	if node, ok := n.Cache["input.NodeId"]; ok {
		var config pbConf.Configuration
		if err := node.DecodeConfig(&config); err != nil {
			log.Fatalln(err)
		}
		
		...
	} else {
		log.Fatalln("...")
    }
}
```

## Testing
We mostly use the same unit test initiation for each plugin. The following code can be used to bootstrap your unit test writing:

```go
package main

import (
	"context"
	"testing"
	testingLib "github.com/MyWeHub/plugin-sdk/testing"
	"github.com/MyWeHub/plugin-sdk/telemetry"
	pb "github.com/MyWeHub/plugin-sdk/gen/pluginrunner"
)

var client pb.PluginRunnerServiceClient

func init() {
	ctx := context.Background()

	t := telemetry.NewTelemetry()
	defer t.ShutdownTracer(ctx)
	defer t.SyncLogger()

	tt := testingLib.NewTesting(t, &serviceServer{})

	client = tt.NewClient(ctx)
}

func TestRunTestV2(t *testing.T) {
	ctx := testingLib.AppendInterceptorTestToken(context.Background())

	_, err := client.RunTestv2(ctx, &pb.InputTestRequestV2{})
	if err != nil {
		t.Fatal(err)
	}
}
```
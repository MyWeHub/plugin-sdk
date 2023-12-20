#Plugin SDK
This library consists of all the code and tools needed to create a plugin in WeHub Fusion Platform and will remove the need to initialize any part of the server and unify all of the plugins initiation code with one single source, so that developers can simply just focus on the logic of each app.

#Address
https://github.com/ariakwehub/plugin_sdk

#Features
- GRPC Server
- HTTP Server
- Telemetry
- Nats
- Connection Service
- Testing [WIP]
- Utilities [WIP]

#Getting Started
Run the command below in your project directory to get the plugin SDK and add it to your project modules:
```bash
go get -u github.com/MyWeHub/plugin-sdk@latest
``` 
The code below will initialize a simple server for plugins.

**main.go**
```go
package main

import (
	"context"
	sdk "github.com/MyWeHub/plugin-sdk"
	"github.com/MyWeHub/plugin-sdk/nats"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	pbconf "workflowplugin/gen/configuration"
)

var (
	logger *zap.Logger
	tracer trace.Tracer
)

func main() {
	ctx := context.Background()

	// telemetry
	t := sdk.NewTelemetry()
	defer t.ShutdownTracer(ctx)
	defer t.SyncLogger()

	logger = t.GetLogger()
	tracer = t.GetTracer()

	//nats
	n := nats.New(&confPB.Configuration{})
	defer n.Close()
	n.Listen(ctx)

	// server
	server := sdk.NewServer()
	server.SetNewGRPC()
	server.RegisterServer(n, newService(), &confPB.Configuration{})
	server.Serve()
}
```

You can also create a custom cache updater function and set it for nats:

```go
package main

import (
   "context"
   "github.com/MyWeHub/plugin-sdk/nats"
   pbconf "workflowplugin/gen/configuration"
)

type natsUpdater struct{}

func (nu *natsUpdater) UpdateCache(configs *[]nats.NodeConfig, cache map[string]*nats.NodeConfig) {
	// Logic here ...
}

func func main() {
    ctx := context.Background()

    //nats
    n := nats.New(&confPB.Configuration{}, &natsUpdater{})
    defer n.Close()
    n.Listen(ctx)
}
```

**server.go**

```go
package main

import (
	"context"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
	pbconf "workflowplugin/gen/configuration"
)

type service struct{}

func newService() *service {
	return &service{}
}

func (s *service) Process(ctx context.Context, in *structpb.Struct, conf proto.Message, action int32, workflowData string) (*structpb.Struct, error) {
        if in == nil || in.Fields == nil {
		return nil, errors.New("input is empty")
	}

	config := conf.(*pbconf.Configuration)

        // Implement Logic...

	return nil, nil
}
```
1. Create a global `logger` and `tracer` instance.
2. Create a service struct. It can consist of any extra field that you need depending on the logic of the application and populate it at runtime, etc...
   2.1. Create a builder function for the service struct. It is better to do so if you aim to populate it [optional]
3. The server accepts a `Process()` method on the service instance passed to the Server Registrar. This is where you will need to implement your application logic.
4. In the main function, create a telemetry instance. This instance is meant to be used both by the library and the application simultaneously. When you create this instance, the library uses your instance as well, so please don't create a second instance of it. Then populate the global variables created in step 1 and use them throughout your application logic for instrumentation.
5. Create a new `Nats` instance, give the pointer to an instance of the generated proto.Configuration of your plugin, and call the `Listen()` method for it to start listening on your plugin-specific inbound configuration in the published workflows. If it receives any configuration, it will decode and cache the value for later use.
6. Configure and start the server by creating a new `server` instance, set a new GRPC server, register your server and pass the nats instance to it, and serve.
   6.1 `server.SetNewGRPC()` accepts an optional `GRPCOptions{}` struct in which allows you to customize your server such as controlling which middlewares to include or change the maximum size of send or receive for each message.
   6.2 There are multiple methods supported on the `server` instance where you can use to further customize your server, but for them to take effect, you must call them before the `server.SetNewGRPC()` method since after this method is called, all of the server properties are set and cannot be changed afterwards.
   6.3 `server.Serve()` accepts an optional`ServerOptions{}` struct in which allows you to declare if you want the library not to start the HTTP server, or to not include graceful shutdown in the program.
   #Environment Variables
   The library needs this set of environment variables in order to function properly (please replace `plugin-name` with your plugin name):
```json
{
       "GRPC_PORT": "6852",
       "HTTP_PORT": "3000",
       "JAEGER_SERVICE_NAME": "plugin-name",
       "NATS": "nats://localhost:4222",
       "OTEL_EXPORTER_JAEGER_AGENT_HOST": "localhost",
       "OTEL_EXPORTER_JAEGER_AGENT_PORT": "6831",
       "OTEL_EXPORTER_JAEGER_ENDPOINT": "http://localhost:14268/api/traces",
       "OTEL_JAEGER_SERVICE_NAME": "plugin-name",
       "PLUGIN_NAME": "plugin-name"
}
```
#Connection Service
This library supports convenient helpers to use connection service simply by creating an instance of it's package and call your desired API:
```go
package main

import (
	"context"
	"github.com/MyWeHub/plugin-sdk/connectionService"
	"log"
)

func main() {
        ctx := context.Background()

        // connection service
	ncs, err := connectionService.New(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer ncs.Close()

	connection, err := ncs.GetConnection(ctx, "63464297-8f51-4094-96be-de25f9b44183")
	if err != nil {
		log.Fatal(err)
	}

	mongoConnection, err := connection.ToMongo()
	if err != nil {
		log.Fatal(err)
	}
}
```
1. Create an instance of the connection service
   1.1 `connectionService.New()` accepts an optional`Options{}` struct in which allows you to declare if you are requesting through your local machine (external request) or not.
2. Call your desired API with the proper parameters (note that the service connection requires the context you pass to include all the proper authentication and module name to fullfil your request).
3. The response of the connection service API supports methods which you can use to convert the `ConnectionMessage` response of it to any supported connection type that you need.

#Testing
In order to use the testing package to simplify the process of creating unit tests, please use the code below:

**server_test.go**

```go
package main

import (
	"context"
	"fmt"
	pb "github.com/MyWeHub/plugin-sdk/gen/pluginrunner"
	testingLib "github.com/MyWeHub/plugin-sdk/testing"
	"testing"
)

var (
	ctx = context.Background()
	client pb.PluginRunnerServiceClient
)

func init() {
	client = testingLib.New(ctx, newService(), &confPB.Configuration{}).NewClient(ctx)
}

func TestRunTestV2(t *testing.T) {
	res, err := c.RunTestv2(ctx, &pb.InputTestRequestV2{
		Inputs:        nil,
		Configuration: nil,
		Action:        0,
		NodeId:        nil,
		Events:        nil,
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("RES")
	fmt.Println(res)
}
```
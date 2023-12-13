package main2

import (
	"context"
	pb "dev.azure.com/WeConnectTechnology/ExchangeHub/_git/wehublib.git/gen/pluginrunner"
	"dev.azure.com/WeConnectTechnology/ExchangeHub/_git/wehublib.git/gen/schema"
	testingLib "dev.azure.com/WeConnectTechnology/ExchangeHub/_git/wehublib.git/testing"
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"reflect"
	"testing"
)

var client pb.PluginRunnerServiceClient

/*func init() {
	ctx := context.Background()

	t := telemetry.NewTelemetry()
	defer t.ShutdownTracer(ctx)
	defer t.SyncLogger()

	tt := testingLib.NewTesting(t, &serviceServer{})

	client = tt.NewClient(ctx)
}*/

func TestTest(t *testing.T) {
	ctx := testingLib.AppendInterceptorTestToken(context.Background())

	_, err := client.RunTest(ctx, &pb.InputTestRequest{})
	if err != nil {
		t.Fatal(err)
	}
}

type Structy struct {
	One   string
	Two   int
	Three bool
}

func TestTMP(t *testing.T) {
	s := &schema.Schema{
		Id:           "ID",
		Name:         "NAME",
		Description:  "DESCRIPTION",
		ParentGroup:  "PARENT_GROUP",
		ClientId:     "CLIENT_ID",
		PartitionKey: "PARTITION_KEY",
		ResourceType: "RESOURCE_TYPE",
		/*One:   "SFDFDSFD",
		Two:   1,
		Three: true,*/
	}

	marshal, err := protojson.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}

	ttt(marshal, &schema.Schema{})
}

func ttt(marshal []byte, m proto.Message) {
	err := protojson.Unmarshal(marshal, m)
	if err != nil {
		panic(err)
	}

	fmt.Println(reflect.TypeOf(m))
	fmt.Println(m.(*schema.Schema))
}

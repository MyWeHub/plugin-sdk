package main2

import (
	"context"
	pb "github.com/MyWeHub/plugin_sdk/gen/pluginrunner"
	"github.com/MyWeHub/plugin_sdk/gen/schema"
	testingLib "github.com/MyWeHub/plugin_sdk/testing"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
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

	_, err := protojson.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}

}

func TestTMP2(t *testing.T) {

	bytes, err := protojson.Marshal(nil)
	if err != nil {
		t.Fatal(err)
	}

	config2 := &anypb.Any{}
	byteConfig := &wrappers.BytesValue{Value: bytes}
	if err := anypb.MarshalFrom(config2, byteConfig, proto.MarshalOptions{}); err != nil {
		t.Fatal(err)
	}

	ttt(config2)
}

func ttt(m proto.Message) {
	var anyout wrappers.BytesValue

	err := anypb.UnmarshalTo(m.(*anypb.Any), &anyout, proto.UnmarshalOptions{})
	if err != nil {
		panic(err)
	}

	err = protojson.Unmarshal(anyout.Value, nil)
	if err != nil {
		panic(err)
	}
}

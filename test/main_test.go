package main2

import (
	"context"
	"testing"
	pb "wehublib/gen/pluginrunner"
	"wehublib/telemetry"
	testingLib "wehublib/testing"
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

func TestTest(t *testing.T) {
	ctx := testingLib.AppendInterceptorTestToken(context.Background())

	_, err := client.RunTest(ctx, &pb.InputTestRequest{})
	if err != nil {
		t.Fatal(err)
	}
}

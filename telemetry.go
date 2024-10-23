package wehublib

import (
	"context"
	apimanagement "github.com/MyWeHub/plugin-sdk/apiManagement"
	"github.com/MyWeHub/plugin-sdk/connectionService"
	"github.com/MyWeHub/plugin-sdk/nats"
	"github.com/MyWeHub/plugin-sdk/organization"
	"github.com/MyWeHub/plugin-sdk/util"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	traceSDK "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"log"
)

type Telemetry struct {
	logger *zap.Logger
	tp     *traceSDK.TracerProvider
	tracer trace.Tracer
}

func NewTelemetry() *Telemetry {
	serviceName := util.GetEnv("PLUGIN_NAME", false, "PLUGIN_NAME", false)
	if serviceName == "" {
		log.Println("WARNING: env 'PLUGIN_NAME' not found")
	}
	tp := newTracerProvider(serviceName)

	logger = newZapLogger()
	tracer = tp.Tracer(serviceName)

	nats.SetTelemetry(logger, tracer)
	connectionService.SetTelemetry(logger, tracer)
	apimanagement.SetTelemetry(logger, tracer)
	organization.SetTelemetry(logger, tracer)

	return &Telemetry{
		logger: logger,
		tp:     tp,
		tracer: tracer,
	}
}

func (t *Telemetry) GetLogger() *zap.Logger {
	return t.logger
}

func (t *Telemetry) GetTracer() trace.Tracer {
	return t.tracer
}

func (t *Telemetry) ShutdownTracer(ctx context.Context) error {
	return t.tp.Shutdown(ctx)
}

func (t *Telemetry) SyncLogger() error {
	return t.logger.Sync()
}

func newTracerProvider(serviceName string) *traceSDK.TracerProvider {
	exporter, err := stdouttrace.New()
	if err != nil {
		panic(err)
	}

	res := resource.NewWithAttributes(semconv.SchemaURL, semconv.ServiceNameKey.String(serviceName))
	tp := traceSDK.NewTracerProvider(traceSDK.WithSampler(traceSDK.AlwaysSample()), traceSDK.WithBatcher(exporter), traceSDK.WithResource(res))

	otel.SetTracerProvider(tp)
	propagator := b3.New(b3.WithInjectEncoding(b3.B3MultipleHeader | b3.B3SingleHeader))
	otel.SetTextMapPropagator(propagator)

	return tp
}

func newZapLogger() *zap.Logger { // TODO: developement/ production option
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	return logger
}

package telemetry

import (
	"context"
	"dev.azure.com/WeConnectTechnology/ExchangeHub/_git/wehublib.git/util"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	traceSDK "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type Telemetry struct {
	logger *zap.Logger
	tp     *traceSDK.TracerProvider
	tracer trace.Tracer
}

func New() *Telemetry {
	serviceName := util.GetEnv("OTEL_JAEGER_SERVICE_NAME", false, "", true)
	tp := newTracerProvider(serviceName)

	return &Telemetry{
		logger: newZapLogger(),
		tp:     tp,
		tracer: tp.Tracer(serviceName),
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
	exporter, err := jaeger.New(jaeger.WithAgentEndpoint())
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

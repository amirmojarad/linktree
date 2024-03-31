package tracing

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"linktree/config"
)

func SetTracerProvider(tracer *trace.TracerProvider) {
	otel.SetTracerProvider(tracer)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
}

func GetTracerProvider(cfg *config.AppConfig) (*trace.TracerProvider, error) {
	return getJaegerTracerProvider(cfg)
}

func getJaegerTracerProvider(cfg *config.AppConfig) (*trace.TracerProvider, error) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint())
	if err != nil {
		return nil, err
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(cfg.App.Name),
			semconv.DeploymentEnvironmentKey.String(cfg.App.Env),
		)),
	)

	return tp, nil
}

package tracer

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

var (
	tp     *tracesdk.TracerProvider //nolint:gochecknoglobals
	tracer trace.Tracer             //nolint:gochecknoglobals
)

func Init(url, name, environment string) error {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return err
	}

	tp = tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(name),
			attribute.String("environment", environment),
		)),
	)
	tracer = tp.Tracer(name)

	return nil
}

func Start(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	if tp == nil {
		return ctx, nil
	}

	return tracer.Start(ctx, spanName, opts...)
}

func End(span trace.Span) {
	if tp == nil || span == nil {
		return
	}

	span.End()
}

func Error(span trace.Span, err error) {
	if tp == nil {
		return
	}

	span.SetStatus(codes.Error, "")
	span.RecordError(err)
}

package tracing

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
)

func InitTracer() *tracesdk.TracerProvider {
    exporter, _ := otlptracegrpc.New(context.Background())
    res, _ := resource.New(context.Background(),
        resource.WithAttributes(semconv.ServiceNameKey.String("main service")),
    )
    otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
        propagation.TraceContext{}, 
    ))

    tp := tracesdk.NewTracerProvider(
        tracesdk.WithBatcher(exporter),
        tracesdk.WithResource(res),
    )

    otel.SetTracerProvider(tp)

    return tp
}
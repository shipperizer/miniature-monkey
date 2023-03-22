package tracing

import (
	"context"
	"runtime/debug"

	"github.com/shipperizer/miniature-monkey/v2/logging"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.18.0"
	"go.opentelemetry.io/otel/trace"
)

type TracerConfig struct {
	ServiceName string
	Endpoint    string
	Logger      logging.LoggerInterface
}

func NewTracerConfig(serviceName, endpoint string, logger logging.LoggerInterface) *TracerConfig {
	c := new(TracerConfig)

	c.ServiceName = serviceName
	c.Endpoint = endpoint
	c.Logger = logger

	return c
}

type Tracer struct {
	tracer trace.Tracer

	logger logging.LoggerInterface
}

func (t *Tracer) init(service string, e sdktrace.SpanExporter) {
	traceProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(e),
		sdktrace.WithResource(
			t.buildResource(service),
		),
	)

	otel.SetTracerProvider(traceProvider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	t.tracer = otel.Tracer(service)
}

func (t *Tracer) gitRevision(settings []debug.BuildSetting) string {
	for _, setting := range settings {
		if setting.Key == "vcs.revision" {
			return setting.Value
		}
	}

	return "n/a"
}

func (t *Tracer) buildResource(service string) *resource.Resource {
	var res *resource.Resource

	res = resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName(service),
		semconv.ServiceVersion("n/a"),
		attribute.String("library", "github.com/shipperizer/miniature-monkey/v2"),
	)

	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		if service == "" {
			service = buildInfo.Path
		}

		res = resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(service),
			semconv.ServiceVersion(t.gitRevision(buildInfo.Settings)),
			attribute.String("app", buildInfo.Main.Path),
		)
	}

	r, _ := resource.Merge(
		resource.Default(),
		res,
	)

	return r
}

func (t *Tracer) Start(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	return t.tracer.Start(ctx, spanName, opts...)
}

// basic tracer implementation of trace.Tracer, just adding some extra configuration
func NewTracer(cfg *TracerConfig) *Tracer {
	t := new(Tracer)

	t.logger = cfg.Logger

	var err error
	var exporter sdktrace.SpanExporter

	// create jaeger exporter
	if cfg.Endpoint != "" {
		exporter, err = jaeger.New(
			jaeger.WithCollectorEndpoint(
				jaeger.WithEndpoint(cfg.Endpoint),
			),
		)
	} else {
		exporter, err = stdouttrace.New(
			stdouttrace.WithPrettyPrint(),
		)
	}

	if err != nil {
		t.logger.Errorf("unable to initialize tracing exporter due: %w", err)
		return nil
	}

	// set tracer provider and propagator properly, this is to ensure all
	// instrumentation library could run well
	t.init(cfg.ServiceName, exporter)

	return t
}

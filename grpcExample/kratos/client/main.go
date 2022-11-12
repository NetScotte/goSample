package main

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	pb "github.com/netscotte/goSample/grpcExample/kratos/proto"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"os"

	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

var (
	ServerHTTPAddress = "localhost:8080"
	Name              = "client"
	JaegerAddress     = "http://localhost:14268"
)

func setTracerProvider(url string) error {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return err
	}
	tp := tracesdk.NewTracerProvider(
		// Set the sampling rate based on the parent span to 100%
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(1.0))),
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in an Resource.
		tracesdk.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(Name),
			attribute.String("env", "dev"),
		)),
	)
	otel.SetTracerProvider(tp)
	return nil
}

func main() {
	logger := log.NewStdLogger(os.Stdout)
	l := log.NewHelper(logger)

	err := setTracerProvider(fmt.Sprintf("%v/api/traces", JaegerAddress))
	if err != nil {
		l.Error(err)
	}
	ctx := context.Background()
	con, err := http.NewClient(
		context.Background(),
		http.WithEndpoint(ServerHTTPAddress),
		http.WithMiddleware(
			tracing.Client(),
		),
	)
	if err != nil {
		l.Error(err)
	}
	client := pb.NewHelloServiceHTTPClient(con)
	reply, err := client.SayHello(ctx, &pb.Request{Name: "goer"})
	if err != nil {
		l.Error(err)
	}
	log.Infof("Receive server response: %v", reply.Message)
}

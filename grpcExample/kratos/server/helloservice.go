package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"io"
	"os"

	pb "github.com/netscotte/goSample/grpcExample/kratos/proto"
)

var (
	Name          = "Hello"
	GRPCAddress   = "localhost:9091"
	HTTPAddress   = "localhost:8081"
	JaegerAddress = "http://localhost:14268"
)

type HelloServiceService struct {
	pb.UnimplementedHelloServiceServer

	log *log.Helper
}

func NewHelloServiceService(l log.Logger) *HelloServiceService {
	return &HelloServiceService{
		log: log.NewHelper(log.With(l, "server", Name)),
	}
}

func GetError() error {
	return fmt.Errorf("%v", "普通错误")
}

func NotFound() error {
	err := GetError()
	return pb.ErrorNotFound(err.Error())
}

func (s *HelloServiceService) SayHello(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	s.log.Infof("receive %v", req.Name)
	msg := fmt.Sprintf("Hello %v", req.Name)
	s.log.Infof("status: %v", req.GetStatus().String())
	s.log.Infof("status: %v", req.GetStatus().Number())
	s.log.Infof("status: %v", req.GetStatus().Enum())
	if req.Name == "error" {
		err := fmt.Errorf("%v", req.Name)
		s.log.Errorf("%v", err)
		return &pb.Response{}, err
	} else if req.Name == "params" {
		err := pb.ErrorParamsError("%v", req.Name)
		s.log.Errorf("%v", err)
		return &pb.Response{}, err
	} else if req.Name == "found" {
		err := NotFound()
		s.log.Errorf("%v", err)
		return &pb.Response{}, err
	}
	s.log.Infof("send %v", msg)
	return &pb.Response{Message: msg}, nil
}

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

type stdLogger struct {
}

func main() {
	logger := NewLogrusLogger()
	logger = log.With(logger, "trace_id", tracing.TraceID())
	logger = log.With(logger, "span_id", tracing.SpanID())
	l := log.NewHelper(logger)

	err := setTracerProvider(fmt.Sprintf("%v/api/traces", JaegerAddress))
	if err != nil {
		l.Error(err)
	}

	httpSrv := http.NewServer(
		http.Address(HTTPAddress),
		http.TLSConfig(GetTlSConfig()),
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
		),
	)
	grpcSrv := grpc.NewServer(
		grpc.Address(GRPCAddress),
		grpc.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
		),
	)

	s := NewHelloServiceService(logger)
	pb.RegisterHelloServiceServer(grpcSrv, s)
	pb.RegisterHelloServiceHTTPServer(httpSrv, s)

	app := kratos.New(
		kratos.Name(Name),
		kratos.Logger(logger),
		kratos.Server(
			grpcSrv,
			httpSrv,
		),
	)
	l.Info("start server")
	if err := app.Run(); err != nil {
		l.Error(err)
	}
}

func GetTlSConfig() *tls.Config {
	certPath := "/Users/netscotte/cert/no_ca_cert/easy.crt"
	certKey := "/Users/netscotte/cert/no_ca_cert/easy.key"
	fp, err := os.Open(certPath)
	if err != nil {
		log.Error(err)
		return nil
	}
	certByte, err := io.ReadAll(fp)
	if err != nil {
		log.Error(err)
		return nil
	}

	fp, err = os.Open(certKey)
	if err != nil {
		log.Error(err)
		return nil
	}
	keyByte, err := io.ReadAll(fp)
	if err != nil {
		log.Error(err)
		return nil
	}

	cert, err := tls.X509KeyPair(certByte, keyByte)
	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}
	return cfg
}

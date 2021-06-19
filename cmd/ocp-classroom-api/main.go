package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"

	"google.golang.org/grpc"

	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-lib/metrics"

	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"

	"github.com/ozoncp/ocp-classroom-api/internal/api"
	prom "github.com/ozoncp/ocp-classroom-api/internal/metrics"
	"github.com/ozoncp/ocp-classroom-api/internal/producer"
	"github.com/ozoncp/ocp-classroom-api/internal/repo"

	desc "github.com/ozoncp/ocp-classroom-api/pkg/ocp-classroom-api"
)

const logPrefix = "ocp-classroom-api service: "

var (
	grpcEndpoint = flag.String("grpc-server-endpoint", "0.0.0.0:7002", "gRPC server endpoint")
	repoArgs     = repo.NewRepoArgs()
	kafkaBroker  = flag.String("kafka-broker", producer.KafkaBroker, "Kafka Apache broker endpoint")
)

func main() {

	flag.Parse()

	introduce()

	log.Debug().Msg(logPrefix + "started")

	initTracing()

	go runMetrics()

	run()
}

func introduce() {
	fmt.Println("Hello World! I'm ocp-classroom-api service by Aleksandr Kuzminykh.")
}

func run() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	listen, err := net.Listen("tcp", *grpcEndpoint)
	if err != nil {
		log.Fatal().Err(err).Msg(logPrefix + "failed to listen")
	}

	repo, err := repo.GetConnectedRepo(ctx, repoArgs)
	if err != nil {
		log.Fatal().Err(err).Msg(logPrefix + "failed to connect to repo")
	}

	logProducer, err := producer.New(ctx, *kafkaBroker)
	if err != nil {
		log.Fatal().Err(err).Msg(logPrefix + "failed to create log producer")
	}

	s := grpc.NewServer()
	desc.RegisterOcpClassroomApiServer(s, api.NewOcpClassroomApi(repo, logProducer))

	log.Debug().Str("endpoint", *grpcEndpoint).Msg(logPrefix + "is listening")

	if err := s.Serve(listen); err != nil {
		log.Fatal().Err(err).Msg(logPrefix + "failed to serve")
	}
}

func runMetrics() {

	prom.RegisterMetrics()
	http.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(":9100", nil)
	if err != nil {
		panic(err)
	}
}

func initTracing() {

	if err := InitGlobalTracer("ocp-classroom-api"); err != nil {
		log.Fatal().Err(err).Msg(logPrefix + "failed to init tracer")
	}
}

// Tip for me: run jaeger in docker and watch trace of MultiCreateClassroomV1 rpc
// at Jaeger UI http://localhost:16686/search

func InitGlobalTracer(serviceName string) error {
	// Sample configuration for testing. Use constant sampling to sample every trace
	// and enable LogSpan to log every span via configured Logger.
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: os.Getenv("JAEGER_AGENT_HOST") + ":" + os.Getenv("JAEGER_AGENT_PORT"),
		},
	}

	// Example logger and metrics factory. Use github.com/uber/jaeger-client-go/log
	// and github.com/uber/jaeger-lib/metrics respectively to bind to real logging and metrics
	// frameworks.
	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	// Initialize tracer with a logger and a metrics factory
	_, err := cfg.InitGlobalTracer(
		serviceName,
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	if err != nil {
		return err
	}

	return nil
}

// CreateClassroomV1
// curl -X POST -d "tenant_id=5&calendar_id=5" localhost:8080/v1/classrooms

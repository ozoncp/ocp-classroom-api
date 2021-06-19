package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/caarlos0/env/v6"
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

// envDefault for local machine
type config struct {
	Endpoint string `env:"ENDPOINT" envDefault:"0.0.0.0:7002"`

	RepoDatabase string `env:"POSTGRES_DB" envDefault:"postgres"`
	RepoUser     string `env:"POSTGRES_USER" envDefault:"postgres"`
	RepoPassword string `env:"POSTGRES_PASSWORD" envDefault:"postgres"`
	RepoEndpoint string `env:"POSTGRES_ENDPOINT" envDefault:"127.0.0.1:5432"`

	KafkaBroker string `env:"KAFKA_BROKER" envDefault:"127.0.0.1:9094"`

	JaegerEndpoint string `env:"JAEGER_ENDPOINT" envDefault:"127.0.0.1:6831"`
}

func main() {

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal().Err(err).Msg(logPrefix + "failed to read configs")
	}

	introduce()

	log.Debug().Msg(logPrefix + "started")

	if err := initGlobalTracer("ocp-classroom-api", cfg.JaegerEndpoint); err != nil {
		log.Fatal().Err(err).Msg(logPrefix + "failed to init tracer")
	}

	go runMetrics()

	run(&cfg)
}

func introduce() {
	fmt.Println("Hello World! I'm ocp-classroom-api service by Aleksandr Kuzminykh.")
}

func run(cfg *config) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	listen, err := net.Listen("tcp", cfg.Endpoint)
	if err != nil {
		log.Fatal().Err(err).Msg(logPrefix + "failed to listen")
	}

	repoArgs := &repo.RepoArgs{
		User:     cfg.RepoUser,
		Password: cfg.RepoPassword,
		Endpoint: cfg.RepoEndpoint,
		DbName:   cfg.RepoDatabase,
		SslMode:  "disable",
	}

	repo, err := repo.GetConnectedRepo(ctx, repoArgs)
	if err != nil {
		log.Fatal().Err(err).Msg(logPrefix + "failed to connect to repo")
	}

	logProducer, err := producer.New(ctx, cfg.KafkaBroker)
	if err != nil {
		log.Fatal().Err(err).Msg(logPrefix + "failed to create log producer")
	}

	s := grpc.NewServer()
	desc.RegisterOcpClassroomApiServer(s, api.NewOcpClassroomApi(repo, logProducer))

	log.Debug().Str("endpoint", cfg.Endpoint).Msg(logPrefix + "is listening")

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

// Tip for me: run jaeger in docker and watch trace of MultiCreateClassroomV1 rpc
// at Jaeger UI http://localhost:16686/search

func initGlobalTracer(serviceName, jaegerEndpoint string) error {
	// Sample configuration for testing. Use constant sampling to sample every trace
	// and enable LogSpan to log every span via configured Logger.
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: jaegerEndpoint,
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

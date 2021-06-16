package main

import (
	"context"
	"flag"
	"net"

	"github.com/rs/zerolog/log"

	"google.golang.org/grpc"

	"github.com/ozoncp/ocp-classroom-api/internal/api"
	"github.com/ozoncp/ocp-classroom-api/internal/producer"
	"github.com/ozoncp/ocp-classroom-api/internal/utils"

	desc "github.com/ozoncp/ocp-classroom-api/pkg/ocp-classroom-api"
)

const logPrefix = "gRPC server: "

func main() {

	log.Debug().Msg(logPrefix + "started")

	var grpcEndpoint = flag.String("grpc-server-endpoint", "0.0.0.0:7002", "gRPC server endpoint")

	flag.Parse()

	if err := InitGlobalTracer("ocp-classroom-api"); err != nil {
		log.Fatal().Err(err).Msg(logPrefix + "failed to init tracer")
	}

	listen, err := net.Listen("tcp", *grpcEndpoint)
	if err != nil {
		log.Fatal().Err(err).Msg(logPrefix + "failed to listen")
	}

	repo, err := utils.GetConnectedRepo(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msg(logPrefix + "failed to connect to repo")
	}
	logProducer, err := producer.New()
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

// CreateClassroomV1
// curl -X POST -d "tenant_id=5&calendar_id=5" localhost:8080/v1/classrooms

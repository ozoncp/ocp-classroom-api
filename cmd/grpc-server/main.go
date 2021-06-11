package main

import (
	"flag"
	"net"

	"github.com/rs/zerolog/log"

	"github.com/ozoncp/ocp-classroom-api/internal/api"

	desc "github.com/ozoncp/ocp-classroom-api/pkg/ocp-classroom-api"
	"google.golang.org/grpc"
)

func main() {

	var grpcEndpoint = flag.String("grpc-server-endpoint", "0.0.0.0:7002", "gRPC server endpoint")

	flag.Parse()

	log.Debug().Msg("doGrpcServerWork...")

	listen, err := net.Listen("tcp", *grpcEndpoint)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to listen")
	}

	s := grpc.NewServer()
	desc.RegisterOcpClassroomApiServer(s, api.NewOcpClassroomApi())

	log.Info().Str("gRPC server endpoint", *grpcEndpoint).Msg("Server listening")
	if err := s.Serve(listen); err != nil {
		log.Fatal().Err(err).Msg("Failed to serve")
	}
}

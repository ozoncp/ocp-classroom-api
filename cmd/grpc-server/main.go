package main

import (
	"flag"
	"net"

	"github.com/rs/zerolog/log"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"

	"google.golang.org/grpc"

	"github.com/ozoncp/ocp-classroom-api/internal/api"
	"github.com/ozoncp/ocp-classroom-api/internal/repo"

	desc "github.com/ozoncp/ocp-classroom-api/pkg/ocp-classroom-api"
)

func getClassroomRepo() *repo.Repo {
	const dbName = "ozon"
	const address = "postgres://postgres:postgres@localhost:5432/" + dbName + "?sslmode=disable"

	db, err := sqlx.Connect("pgx", address)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to postgres")
	}

	log.Debug().Msgf("Connected to DB %v", dbName)

	classroomRepo := repo.New(db)

	return &classroomRepo
}

func main() {

	var grpcEndpoint = flag.String("grpc-server-endpoint", "0.0.0.0:7002", "gRPC server endpoint")

	flag.Parse()

	log.Debug().Msg("doGrpcServerWork...")

	listen, err := net.Listen("tcp", *grpcEndpoint)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to listen")
	}

	s := grpc.NewServer()
	desc.RegisterOcpClassroomApiServer(s, api.NewOcpClassroomApi(*getClassroomRepo()))

	log.Info().Str("gRPC server endpoint", *grpcEndpoint).Msg("Server listening")
	if err := s.Serve(listen); err != nil {
		log.Fatal().Err(err).Msg("Failed to serve")
	}
}

package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"

	desc "github.com/ozoncp/ocp-classroom-api/pkg/ocp-classroom-api"
	"google.golang.org/grpc"
)

func main() {

	var grpcEndpoint = flag.String("grpc-server-endpoint", "0.0.0.0:7002", "gRPC server endpoint")

	flag.Parse()

	log.Debug().Msg("doGrpcClientWork...")

	conn, err := grpc.Dial(*grpcEndpoint, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect")
	}
	defer conn.Close()

	log.Debug().Str("gRPC server endpoint", *grpcEndpoint).Msg("Client connected")

	c := desc.NewOcpClassroomApiClient(conn)

	var cmd string
	fmt.Print("What to do? ('l' - List, 'c' - Create, 'd' - Describe', 'r' - Remove): ")
	fmt.Scan(&cmd)
	fmt.Println()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	switch cmd {

	case "l":
		r, err := c.ListClassroomsV1(ctx, &desc.ListClassroomsV1Request{Limit: 5, Offset: 0})
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to list classrooms")
		}

		log.Debug().Msgf("Response on ListClassroom %v", r)

	case "c":
		r, err := c.CreateClassroomV1(ctx, &desc.CreateClassroomV1Request{TenantId: 1, CalendarId: 1})
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to create classroom")
		}

		log.Debug().Msgf("Response on CreateClassrooms %v", r)

	case "d":
		r, err := c.DescribeClassroomV1(ctx, &desc.DescribeClassroomV1Request{ClassroomId: 1, Verbose: true})
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to describe classroom")
		}

		log.Debug().Msgf("Response on DescribeClassroom %v", r)

	case "r":
		r, err := c.RemoveClassroomV1(ctx, &desc.RemoveClassroomV1Request{ClassroomId: 3})
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to remove classroom")
		}

		log.Debug().Msgf("Response on RemoveClassroom %v", r)
	}
}

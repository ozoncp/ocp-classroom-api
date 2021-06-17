package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/ozoncp/ocp-classroom-api/internal/models"
	"github.com/ozoncp/ocp-classroom-api/internal/utils"
	desc "github.com/ozoncp/ocp-classroom-api/pkg/ocp-classroom-api"
	"google.golang.org/grpc"
)

const logPrefix = "gRPC client: "

var grpcEndpoint = flag.String("grpc-server-endpoint", "0.0.0.0:7002", "gRPC server endpoint")

func main() {

	flag.Parse()

	log.Debug().Msg(logPrefix + "started")

	conn, err := grpc.Dial(*grpcEndpoint, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatal().Err(err).Msg(logPrefix + "failed to connect")
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error().Err(err).Msg(logPrefix + "failed to close connection")
		}
	}()

	log.Debug().Str("gRPC server endpoint", *grpcEndpoint).Msg(logPrefix + "connected")

	client := desc.NewOcpClassroomApiClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()

	for {

		var cmd string
		fmt.Print("What to do? (",
			"'l' - List,",
			"'c' - Create,",
			"'mc' - MultiCreate,",
			"'d' - Describe',",
			"'u' - Update,",
			"'r' - Remove,",
			"'x' - Exit",
			"): ")
		fmt.Scan(&cmd)

		switch cmd {

		case "l":

			func() {

				var limit uint64
				var offset uint64
				fmt.Print("Enter the limit and offset: ")
				fmt.Scan(&limit, &offset)

				req := &desc.ListClassroomsV1Request{Limit: limit, Offset: offset}
				var res *desc.ListClassroomsV1Response
				var err error

				defer utils.LogGrpcCall(logPrefix+"ListClassroomsV1 call", &req, &res, &err)

				res, err = client.ListClassroomsV1(ctx, req)
				if err != nil {
					return
				}
			}()

		case "c":

			func() {

				var tenantId uint64
				var calendarId uint64
				fmt.Print("Enter tenantId and calendarId: ")
				fmt.Scan(&tenantId, &calendarId)

				req := &desc.CreateClassroomV1Request{TenantId: tenantId, CalendarId: calendarId}
				var res *desc.CreateClassroomV1Response
				var err error

				defer utils.LogGrpcCall(logPrefix+"CreateClassroomV1 call", &req, &res, &err)

				res, err = client.CreateClassroomV1(ctx, req)
			}()

		case "mc":

			func() {

				var count int
				fmt.Print("Enter count: ")
				fmt.Scan(&count)

				if count < 1 {
					log.Fatal().Msg(logPrefix + "count can not be less 1")
				}

				req := &desc.MultiCreateClassroomV1Request{}
				var res *desc.MultiCreateClassroomV1Response
				var err error

				defer utils.LogGrpcCall(logPrefix+"MultiCreateClassroomV1 call", &req, &res, &err)

				for i := 0; i < count; i++ {

					var tenantId uint64
					var calendarId uint64
					fmt.Print("Enter tenantId and calendarId: ")
					fmt.Scan(&tenantId, &calendarId)

					req.Classrooms = append(req.Classrooms,
						&desc.CreateClassroomV1Request{TenantId: tenantId, CalendarId: calendarId})
				}

				res, err = client.MultiCreateClassroomV1(ctx, req)
			}()

		case "d":

			func() {

				var classroomId uint64
				fmt.Print("Enter classroomId: ")
				fmt.Scan(&classroomId)

				req := &desc.DescribeClassroomV1Request{ClassroomId: classroomId}
				var res *desc.DescribeClassroomV1Response
				var err error

				defer utils.LogGrpcCall(logPrefix+"DescribeClassroomV1 call", &req, &res, &err)

				res, err = client.DescribeClassroomV1(ctx, req)
			}()

		case "u":

			func() {

				var classroom models.Classroom
				fmt.Print("Enter classroomId, tenantId and calendarId: ")
				fmt.Scan(&classroom.Id, &classroom.TenantId, &classroom.CalendarId)

				req := &desc.UpdateClassroomV1Request{Classroom: classroom.ToProtoClassroom()}
				var res *desc.UpdateClassroomV1Response
				var err error

				defer utils.LogGrpcCall(logPrefix+"UpdateClassroomV1 call", &req, &res, &err)

				res, err = client.UpdateClassroomV1(ctx, req)
			}()

		case "r":

			func() {

				var classroomId uint64
				fmt.Print("Enter classroomId: ")
				fmt.Scan(&classroomId)

				req := &desc.RemoveClassroomV1Request{ClassroomId: classroomId}
				var res *desc.RemoveClassroomV1Response
				var err error

				defer utils.LogGrpcCall(logPrefix+"RemoveClassroomV1 call", &req, &res, &err)

				res, err = client.RemoveClassroomV1(ctx, req)
			}()

		case "x":

			return
		}
	}
}

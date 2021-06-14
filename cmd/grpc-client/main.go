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
	fmt.Print("What to do? ('l' - List, 'c' - Create, 'mc' - MultiCreate, 'd' - Describe', 'u' - Update, 'r' - Remove): ")
	fmt.Scan(&cmd)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()

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

			defer utils.LogGrpcCall("ListClassroomsV1 call", &req, &res, &err)

			res, err = c.ListClassroomsV1(ctx, req)
			if err != nil {
				return
			}
		}()

	case "c":

		func() {

			var tenant_id uint64
			var calendar_id uint64
			fmt.Print("Enter tenant_id and calendar_id: ")
			fmt.Scan(&tenant_id, &calendar_id)

			req := &desc.CreateClassroomV1Request{TenantId: tenant_id, CalendarId: calendar_id}
			var res *desc.CreateClassroomV1Response
			var err error

			defer utils.LogGrpcCall("CreateClassroomV1 call", &req, &res, &err)

			res, err = c.CreateClassroomV1(ctx, req)
			if err != nil {
				return
			}
		}()

	case "mc":

		func() {

			var count int
			fmt.Print("Enter count: ")
			fmt.Scan(&count)

			if count < 1 {
				log.Fatal().Msg("Count can not be less 1")
			}

			req := &desc.MultiCreateClassroomV1Request{}
			var res *desc.MultiCreateClassroomV1Response
			var err error

			defer utils.LogGrpcCall("MultiCreateClassroomV1 call", &req, &res, &err)

			for i := 0; i < count; i++ {

				var tenant_id uint64
				var calendar_id uint64
				fmt.Print("Enter tenant_id and calendar_id: ")
				fmt.Scan(&tenant_id, &calendar_id)

				req.Classrooms = append(req.Classrooms,
					&desc.CreateClassroomV1Request{TenantId: tenant_id, CalendarId: calendar_id})
			}

			res, err = c.MultiCreateClassroomV1(ctx, req)
			if err != nil {
				return
			}
		}()

	case "d":

		func() {

			var classroom_id uint64
			fmt.Print("Enter classroom_id: ")
			fmt.Scan(&classroom_id)

			req := &desc.DescribeClassroomV1Request{ClassroomId: classroom_id}
			var res *desc.DescribeClassroomV1Response
			var err error

			defer utils.LogGrpcCall("DescribeClassroomV1 call", &req, &res, &err)

			res, err = c.DescribeClassroomV1(ctx, req)
			if err != nil {
				return
			}
		}()

	case "u":

		func() {

			var classroom models.Classroom
			fmt.Print("Enter classroom_id, tenant_id and calendar_id: ")
			fmt.Scan(&classroom.Id, &classroom.TenantId, &classroom.CalendarId)

			req := &desc.UpdateClassroomV1Request{Classroom: classroom.ToProtoClassroom()}
			var res *desc.UpdateClassroomV1Response
			var err error

			defer utils.LogGrpcCall("UpdateClassroomV1 call", &req, &res, &err)

			res, err = c.UpdateClassroomV1(ctx, req)
			if err != nil {
				return
			}
		}()

	case "r":

		func() {

			var classroom_id uint64
			fmt.Print("Enter classroom_id: ")
			fmt.Scan(&classroom_id)

			req := &desc.RemoveClassroomV1Request{ClassroomId: classroom_id}
			var res *desc.RemoveClassroomV1Response
			var err error

			defer utils.LogGrpcCall("RemoveClassroomV1 call", &req, &res, &err)

			res, err = c.RemoveClassroomV1(ctx, req)
			if err != nil {
				return
			}
		}()
	}
}

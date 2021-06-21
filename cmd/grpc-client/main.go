package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"

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

		switch getCommandFromUserInput() {

		case "l":

			func() {

				req := getListRequestFromUserInput()
				var res *desc.ListClassroomsV1Response
				var err error

				defer utils.LogGrpcCall(logPrefix+"ListClassroomsV1 call", &req, &res, &err)
				res, err = client.ListClassroomsV1(ctx, req)
			}()

		case "c":

			func() {

				req := getCreateRequestFromUserInput()
				var res *desc.CreateClassroomV1Response
				var err error

				defer utils.LogGrpcCall(logPrefix+"CreateClassroomV1 call", &req, &res, &err)
				res, err = client.CreateClassroomV1(ctx, req)
			}()

		case "mc":

			func() {

				req := getMultiCreateRequestFromUserInput()
				var res *desc.MultiCreateClassroomV1Response
				var err error

				defer utils.LogGrpcCall(logPrefix+"MultiCreateClassroomV1 call", &req, &res, &err)
				res, err = client.MultiCreateClassroomV1(ctx, req)
			}()

		case "d":

			func() {

				req := getDescribeRequestFromUserInput()
				var res *desc.DescribeClassroomV1Response
				var err error

				defer utils.LogGrpcCall(logPrefix+"DescribeClassroomV1 call", &req, &res, &err)
				res, err = client.DescribeClassroomV1(ctx, req)
			}()

		case "u":

			func() {

				req := getUpdateRequestFromUserInput()
				var res *desc.UpdateClassroomV1Response
				var err error

				defer utils.LogGrpcCall(logPrefix+"UpdateClassroomV1 call", &req, &res, &err)
				res, err = client.UpdateClassroomV1(ctx, req)
			}()

		case "r":

			func() {

				req := getRemoveRequestFromUserInput()
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

func getCommandFromUserInput() (cmd string) {

	for {
		fmt.Print("What to do? (",
			"'l' - List,",
			"'c' - Create,",
			"'mc' - MultiCreate,",
			"'d' - Describe',",
			"'u' - Update,",
			"'r' - Remove,",
			"'x' - Exit",
			"): ")

		if _, err := fmt.Scan(&cmd); err != nil {
			fmt.Println("Error occurred", err, ". Try again")
			continue
		}

		break
	}

	return
}

func getListRequestFromUserInput() *desc.ListClassroomsV1Request {

	var limit, offset uint64
	for {
		fmt.Print("Enter the limit and offset: ")

		if _, err := fmt.Scan(&limit, &offset); err != nil {
			fmt.Println("Error occurred", err, ". Try again")
			continue
		}

		break
	}

	return &desc.ListClassroomsV1Request{Limit: limit, Offset: offset}
}

func getCreateRequestFromUserInput() *desc.CreateClassroomV1Request {

	var tenantId, calendarId uint64
	for {
		fmt.Print("Enter tenantId and calendarId: ")

		if _, err := fmt.Scan(&tenantId, &calendarId); err != nil {
			fmt.Println("Error occurred", err, ". Try again")
			continue
		}

		break
	}

	return &desc.CreateClassroomV1Request{TenantId: tenantId, CalendarId: calendarId}
}

func getMultiCreateRequestFromUserInput() *desc.MultiCreateClassroomV1Request {

	var count int
	for {
		fmt.Print("Enter count: ")

		if _, err := fmt.Scan(&count); err != nil {
			fmt.Println("Error occurred", err, ". Try again")
			continue
		}

		if count < 1 {
			fmt.Println("Count can not be less 1. Try again")
			continue
		}

		break
	}

	req := &desc.MultiCreateClassroomV1Request{}

	for i := 0; i < count; i++ {

		var tenantId uint64
		var calendarId uint64
		for {
			fmt.Print("Enter tenantId and calendarId: ")

			if _, err := fmt.Scan(&tenantId, &calendarId); err != nil {
				fmt.Println("Error occurred", err, ". Try again")
				continue
			}

			break
		}

		req.Classrooms = append(req.Classrooms,
			&desc.CreateClassroomV1Request{TenantId: tenantId, CalendarId: calendarId})
	}

	return req
}

func getDescribeRequestFromUserInput() *desc.DescribeClassroomV1Request {

	var classroomId uint64
	for {
		fmt.Print("Enter classroomId: ")

		if _, err := fmt.Scan(&classroomId); err != nil {
			fmt.Println("Error occurred", err, ". Try again")
			continue
		}

		break
	}

	return &desc.DescribeClassroomV1Request{ClassroomId: classroomId}
}

func getUpdateRequestFromUserInput() *desc.UpdateClassroomV1Request {

	var classroomId, tenantId, calendarId uint64
	for {
		fmt.Print("Enter classroomId, tenantId and calendarId: ")

		if _, err := fmt.Scan(&classroomId, &tenantId, &calendarId); err != nil {
			fmt.Println("Error occurred", err, ". Try again")
			continue
		}

		break
	}

	return &desc.UpdateClassroomV1Request{
		Classroom: &desc.Classroom{
			ClassroomId: classroomId,
			TenantId:    tenantId,
			CalendarId:  calendarId,
		},
	}
}

func getRemoveRequestFromUserInput() *desc.RemoveClassroomV1Request {

	var classroomId uint64
	for {
		fmt.Print("Enter classroomId: ")

		if _, err := fmt.Scan(&classroomId); err != nil {
			fmt.Println("Error occurred", err, ". Try again")
			continue
		}

		break
	}

	return &desc.RemoveClassroomV1Request{ClassroomId: classroomId}
}

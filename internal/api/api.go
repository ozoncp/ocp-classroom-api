package api

import (
	"context"

	grpcApi "github.com/ozoncp/ocp-classroom-api/pkg/ocp-classroom-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	errClassroomNotFound        = "classroom not found"
	errClassroomCannotBeCreated = "classroom not created"
)

type api struct {
	grpcApi.UnimplementedOcpClassroomApiServer
}

func (a *api) ListClassroomsV1(ctx context.Context,
	req *grpcApi.ListClassroomsV1Request) (*grpcApi.ListClassroomsV1Response, error) {

	if err := req.Validate(); err != nil {

		log.Err(err).Msg("request can not be validated")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Debug().
		Uint64("Limit", req.Limit).
		Uint64("Offset", req.Offset).
		Msg("ListClassroomV1 call")

	err := status.Error(codes.NotFound, errClassroomNotFound)
	return nil, err
}

func (a *api) DescribeClassroomV1(ctx context.Context,
	req *grpcApi.DescribeClassroomV1Request) (*grpcApi.DescribeClassroomV1Response, error) {

	if err := req.Validate(); err != nil {

		log.Err(err).Msg("request can not be validated")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Debug().
		Uint64("ClassroomId", req.ClassroomId).
		Bool("Verbose", req.Verbose).
		Msg("DescribeClassroomV1 call")

	err := status.Error(codes.NotFound, errClassroomNotFound)
	return nil, err
}

func (a *api) CreateClassroomV1(ctx context.Context,
	req *grpcApi.CreateClassroomV1Request) (*grpcApi.CreateClassroomV1Response, error) {

	if err := req.Validate(); err != nil {

		log.Err(err).Msg("request can not be validated")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Debug().
		Uint64("TenantId", req.TenantId).
		Uint64("CalendarId", req.CalendarId).
		Msg("CreateClassroomV1 call")

	err := status.Error(codes.NotFound, errClassroomCannotBeCreated)
	return nil, err
}

func (a *api) RemoveClassroomV1(ctx context.Context,
	req *grpcApi.RemoveClassroomV1Request) (*grpcApi.RemoveClassroomV1Response, error) {

	if err := req.Validate(); err != nil {

		log.Err(err).Msg("request can not be validated")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Debug().
		Uint64("ClassroomId", req.ClassroomId).
		Msg("RemoveClassroomV1 call")

	err := status.Error(codes.NotFound, errClassroomNotFound)
	return nil, err
}

func NewOcpClassroomApi() grpcApi.OcpClassroomApiServer {
	return &api{}
}

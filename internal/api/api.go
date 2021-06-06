package api

import (
	"context"

	grpcApi "github.com/ozoncp/ocp-classroom-api/pkg/ocp-classroom-api"
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
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := status.Error(codes.NotFound, errClassroomNotFound)
	return nil, err
}

func (a *api) DescribeClassroomV1(ctx context.Context,
	req *grpcApi.DescribeClassroomV1Request) (*grpcApi.DescribeClassroomV1Response, error) {

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := status.Error(codes.NotFound, errClassroomNotFound)
	return nil, err
}

func (a *api) CreateClassroomV1(ctx context.Context,
	req *grpcApi.CreateClassroomV1Request) (*grpcApi.CreateClassroomV1Response, error) {

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := status.Error(codes.NotFound, errClassroomCannotBeCreated)
	return nil, err
}

func (a *api) RemoveClassroomV1(ctx context.Context,
	req *grpcApi.RemoveClassroomV1Request) (*grpcApi.RemoveClassroomV1Response, error) {

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := status.Error(codes.NotFound, errClassroomNotFound)
	return nil, err
}

func NewOcpClassroomApi() grpcApi.OcpClassroomApiServer {
	return &api{}
}

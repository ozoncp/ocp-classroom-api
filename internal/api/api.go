package api

import (
	"context"

	desc "github.com/ozoncp/ocp-classroom-api/pkg/ocp-classroom-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const errClassroomNotFound = "classroom not found"

type api struct {
	desc.UnimplementedOcpClassroomApiServer
}

func (a *api) DescribeClassroomV1(ctx context.Context,
	req *desc.DescribeClassroomV1Request) (*desc.DescribeClassroomV1Response, error) {

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := status.Error(codes.NotFound, errClassroomNotFound)
	return nil, err
}

func NewOcpClassroomApi() desc.OcpClassroomApiServer {
	return &api{}
}

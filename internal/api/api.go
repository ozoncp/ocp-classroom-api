package api

import (
	"context"

	"github.com/rs/zerolog/log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ozoncp/ocp-classroom-api/internal/models"
	"github.com/ozoncp/ocp-classroom-api/internal/repo"
	grpcApi "github.com/ozoncp/ocp-classroom-api/pkg/ocp-classroom-api"
)

type api struct {
	grpcApi.UnimplementedOcpClassroomApiServer
	classroomRepo repo.Repo
}

func (a *api) ListClassroomsV1(ctx context.Context,
	req *grpcApi.ListClassroomsV1Request) (*grpcApi.ListClassroomsV1Response, error) {

	if err := req.Validate(); err != nil {

		log.Error().Err(err).Msg("Request failed validation")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	classrooms, err := a.classroomRepo.ListClassrooms(ctx, req.Limit, req.Offset)
	if err != nil {

		log.Error().Err(err).Msg("Failed to list classrooms")
		return nil, err
	}

	var protoClassrooms []*grpcApi.Classroom

	for _, classroom := range classrooms {

		protoClassrooms = append(protoClassrooms, classroom.ToProtoClassroom())
	}

	log.Debug().
		Uint64("Limit", req.Limit).
		Uint64("Offset", req.Offset).
		Msgf("ListClassroomV1 call: %v", protoClassrooms)

	return &grpcApi.ListClassroomsV1Response{Classrooms: protoClassrooms}, nil
}

func (a *api) DescribeClassroomV1(ctx context.Context,
	req *grpcApi.DescribeClassroomV1Request) (*grpcApi.DescribeClassroomV1Response, error) {

	if err := req.Validate(); err != nil {

		log.Error().Err(err).Msg("Request failed validation")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	classroom, err := a.classroomRepo.DescribeClassroom(ctx, req.ClassroomId)
	if err != nil {

		log.Error().Err(err).Msg("Failed to describe classroom")
		return nil, err
	}

	protoClassroom := classroom.ToProtoClassroom()

	log.Debug().
		Uint64("ClassroomId", req.ClassroomId).
		Bool("Verbose", req.Verbose).
		Msgf("DescribeClassroomV1 call: %v", protoClassroom)

	return &grpcApi.DescribeClassroomV1Response{Classroom: protoClassroom}, nil
}

func (a *api) CreateClassroomV1(ctx context.Context,
	req *grpcApi.CreateClassroomV1Request) (*grpcApi.CreateClassroomV1Response, error) {

	if err := req.Validate(); err != nil {

		log.Error().Err(err).Msg("Request failed validation")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	classroomId, err := a.classroomRepo.AddClassroom(ctx, models.Classroom{
		TenantId:   req.TenantId,
		CalendarId: req.CalendarId,
	})
	if err != nil {

		log.Error().Err(err).Msg("Failed to add classroom")
		return nil, err
	}

	log.Debug().
		Uint64("TenantId", req.TenantId).
		Uint64("CalendarId", req.CalendarId).
		Msgf("CreateClassroomV1 call: %v", classroomId)

	return &grpcApi.CreateClassroomV1Response{ClassroomId: classroomId}, nil
}

func (a *api) RemoveClassroomV1(ctx context.Context,
	req *grpcApi.RemoveClassroomV1Request) (*grpcApi.RemoveClassroomV1Response, error) {

	if err := req.Validate(); err != nil {

		log.Error().Err(err).Msg("Request failed validation")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	found, err := a.classroomRepo.RemoveClassroom(ctx, req.ClassroomId)
	if err != nil {

		log.Error().Err(err).Msg("Failed to remove classroom")
		return nil, err
	}

	log.Debug().
		Uint64("ClassroomId", req.ClassroomId).
		Msgf("RemoveClassroomV1 call: %v", found)

	return &grpcApi.RemoveClassroomV1Response{Found: found}, nil
}

func NewOcpClassroomApi(classroomRepo repo.Repo) grpcApi.OcpClassroomApiServer {
	return &api{classroomRepo: classroomRepo}
}

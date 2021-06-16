package api

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	opentracing "github.com/opentracing/opentracing-go"

	"github.com/ozoncp/ocp-classroom-api/internal/flusher"
	"github.com/ozoncp/ocp-classroom-api/internal/models"
	"github.com/ozoncp/ocp-classroom-api/internal/producer"
	"github.com/ozoncp/ocp-classroom-api/internal/repo"
	"github.com/ozoncp/ocp-classroom-api/internal/utils"
	grpcApi "github.com/ozoncp/ocp-classroom-api/pkg/ocp-classroom-api"
)

// chunkSize is used for MultiCreateClassroomV1
const chunkSize int = 5

// api implements gRPC server and operates with DB
type api struct {
	grpcApi.UnimplementedOcpClassroomApiServer
	classroomRepo repo.Repo
	logProducer   producer.LogProducer
}

// ListClassroomsV1 returns list of classrooms from DB
func (a *api) ListClassroomsV1(ctx context.Context,
	req *grpcApi.ListClassroomsV1Request) (res *grpcApi.ListClassroomsV1Response, err error) {

	defer utils.LogGrpcCall("CreateClassroomV1", &req, &res, &err)

	if err = req.Validate(); err != nil {

		err = status.Error(codes.InvalidArgument, err.Error())
		return nil, err
	}

	classrooms, err := a.classroomRepo.ListClassrooms(ctx, req.Limit, req.Offset)
	if err != nil {

		err = status.Error(codes.Unavailable, err.Error())
		return nil, err
	}

	var protoClassrooms []*grpcApi.Classroom
	for _, classroom := range classrooms {

		protoClassrooms = append(protoClassrooms, classroom.ToProtoClassroom())
	}

	res = &grpcApi.ListClassroomsV1Response{Classrooms: protoClassrooms}
	return res, nil
}

// DescribeClassroomV1 returns classroom from DB requested by id
func (a *api) DescribeClassroomV1(ctx context.Context,
	req *grpcApi.DescribeClassroomV1Request) (res *grpcApi.DescribeClassroomV1Response, err error) {

	defer utils.LogGrpcCall("CreateClassroomV1", &req, &res, &err)

	if err = req.Validate(); err != nil {

		err = status.Error(codes.InvalidArgument, err.Error())
		return nil, err
	}

	classroom, err := a.classroomRepo.DescribeClassroom(ctx, req.ClassroomId)
	if err != nil {

		err = status.Error(codes.Unavailable, err.Error())
		return nil, err
	}

	protoClassroom := classroom.ToProtoClassroom()

	res = &grpcApi.DescribeClassroomV1Response{Classroom: protoClassroom}
	return res, nil
}

// CreateClassroomV1 returns id of created classroom in DB by passed tenant_id and calendar_id
func (a *api) CreateClassroomV1(ctx context.Context,
	req *grpcApi.CreateClassroomV1Request) (res *grpcApi.CreateClassroomV1Response, err error) {

	defer utils.LogGrpcCall("CreateClassroomV1", &req, &res, &err)
	defer func() {

		if errKafka := a.logProducer.Send(producer.Create, req, res, err); errKafka != nil {
			err = errKafka
		}
	}()

	if err = req.Validate(); err != nil {

		err = status.Error(codes.InvalidArgument, err.Error())
		return nil, err
	}

	classroomId, err := a.classroomRepo.AddClassroom(ctx, models.Classroom{
		TenantId:   req.TenantId,
		CalendarId: req.CalendarId,
	})
	if err != nil {

		err = status.Error(codes.Unavailable, err.Error())
		return nil, err
	}

	res = &grpcApi.CreateClassroomV1Response{ClassroomId: classroomId}
	return res, nil
}

// MultiCreateClassroomV1 returns count of created classrooms in DB by passed list of tenant_id and calendar_id
func (a *api) MultiCreateClassroomV1(ctx context.Context,
	req *grpcApi.MultiCreateClassroomV1Request) (res *grpcApi.MultiCreateClassroomV1Response, err error) {

	defer utils.LogGrpcCall("MultiCreateClassroomV1", &req, &res, &err)

	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("MultiCreateClassroomV1")
	defer span.Finish()

	if err = req.Validate(); err != nil {

		err = status.Error(codes.InvalidArgument, err.Error())
		return nil, err
	}

	var classrooms []models.Classroom
	for _, protoClassroom := range req.Classrooms {

		classrooms = append(classrooms, models.Classroom{
			TenantId:   protoClassroom.TenantId,
			CalendarId: protoClassroom.CalendarId,
		})
	}

	fl := flusher.New(a.classroomRepo, chunkSize)
	remainingClassrooms := fl.Flush(ctx, span, classrooms)

	var createdCount = uint64(len(classrooms) - len(remainingClassrooms))
	if createdCount == 0 {

		err = status.Error(codes.Unavailable, errors.New("flush call returned non nil result").Error())
		return nil, err
	}

	res = &grpcApi.MultiCreateClassroomV1Response{CreatedCount: createdCount}
	return res, nil
}

// UpdateClassroomV1 changes classroom in DB by passed id and new tenant_id and calendar_id
// and returns whether classroom was changed
func (a *api) UpdateClassroomV1(ctx context.Context,
	req *grpcApi.UpdateClassroomV1Request) (res *grpcApi.UpdateClassroomV1Response, err error) {

	defer utils.LogGrpcCall("UpdateClassroomV1", &req, &res, &err)
	defer func() {

		if errKafka := a.logProducer.Send(producer.Update, req, res, err); errKafka != nil {
			err = errKafka
		}
	}()

	if err = req.Validate(); err != nil {

		err = status.Error(codes.InvalidArgument, err.Error())
		return nil, err
	}

	classroom := models.FromProtoClassroom(req.Classroom)

	found, err := a.classroomRepo.UpdateClassroom(ctx, *classroom)
	if err != nil {

		err = status.Error(codes.Unavailable, err.Error())
		return nil, err
	}

	res = &grpcApi.UpdateClassroomV1Response{Found: found}
	return res, nil
}

// RemoveClassroomV1 removes classroom from DB by passed id and returs whether classroom was removed
func (a *api) RemoveClassroomV1(ctx context.Context,
	req *grpcApi.RemoveClassroomV1Request) (res *grpcApi.RemoveClassroomV1Response, err error) {

	defer utils.LogGrpcCall("RemoveClassroomV1", &req, &res, &err)
	defer func() {

		if errKafka := a.logProducer.Send(producer.Remove, req, res, err); errKafka != nil {
			err = errKafka
		}
	}()

	if err = req.Validate(); err != nil {

		err = status.Error(codes.InvalidArgument, err.Error())
		return nil, err
	}

	found, err := a.classroomRepo.RemoveClassroom(ctx, req.ClassroomId)
	if err != nil {

		err = status.Error(codes.Unavailable, err.Error())
		return nil, err
	}

	res = &grpcApi.RemoveClassroomV1Response{Found: found}
	return res, nil
}

// NewOcpClassroomApi returns implementation of OcpClassroomApiServer interface to operate DB
func NewOcpClassroomApi(
	classroomRepo repo.Repo,
	logProducer producer.LogProducer) grpcApi.OcpClassroomApiServer {

	return &api{classroomRepo: classroomRepo, logProducer: logProducer}
}

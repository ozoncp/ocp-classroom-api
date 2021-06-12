package repo

import (
	"context"

	"github.com/ozoncp/ocp-classroom-api/internal/models"
)

type Repo interface {
	ListClassrooms(ctx context.Context, limit, offset uint64) ([]models.Classroom, error)
	AddClassrooms(ctx context.Context, classrooms []models.Classroom) error
	RemoveClassroom(ctx context.Context, classroomId uint64) error
	DescribeClassroom(ctx context.Context, classroomId uint64) (*models.Classroom, error)
}

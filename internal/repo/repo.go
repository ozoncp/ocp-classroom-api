package repo

import "github.com/ozoncp/ocp-classroom-api/internal/models"

type Repo interface {
	AddClassrooms(classrooms []models.Classroom) error
	RemoveClassroom(classroomId uint64) error
	DescribeClassroom(classroomId uint64) (*models.Classroom, error)
	ListClassrooms(limit, offset uint64) ([]models.Classroom, error)
}

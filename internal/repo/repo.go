package repo

import "github.com/ozoncp/ocp-classroom-api/internal/models"

type Repo interface {
	AddTasks(task []models.Classroom) error
	RemoveTask(taskId uint64) error
	DescribeTask(taskId uint64) (*models.Classroom, error)
	ListTasks(limit, offset uint64) ([]models.Classroom, error)
}

package flusher

import (
	"github.com/ozoncp/ocp-classroom-api/internal/models"
	"github.com/ozoncp/ocp-classroom-api/internal/repo"
)

type Flusher interface {
	Flush(classrooms []models.Classroom) []models.Classroom
}

type flusher struct {
	repo repo.Repo
}

func New(repo repo.Repo) *flusher {
	return &flusher{repo: repo}
}

func (fl *flusher) Flush(classrooms []models.Classroom) []models.Classroom {

	if err := fl.repo.AddClassrooms(classrooms); err != nil {
		return classrooms
	}

	return nil
}

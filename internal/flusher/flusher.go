package flusher

import (
	"github.com/ozoncp/ocp-classroom-api/internal/models"
	"github.com/ozoncp/ocp-classroom-api/internal/repo"
	"github.com/ozoncp/ocp-classroom-api/internal/utils"
)

type Flusher interface {
	Flush(classrooms []models.Classroom) []models.Classroom
}

type flusher struct {
	repo      repo.Repo
	chunkSize int
}

func New(repo repo.Repo, chunkSize int) *flusher {
	return &flusher{repo: repo, chunkSize: chunkSize}
}

func (fl *flusher) Flush(classrooms []models.Classroom) []models.Classroom {

	chunks, err := utils.SplitSlice(classrooms, fl.chunkSize)

	if err != nil {
		return classrooms
	}

	for i, chunk := range chunks {

		if err := fl.repo.AddClassrooms(chunk); err != nil {
			return classrooms[fl.chunkSize*i:]
		}
	}

	return nil
}

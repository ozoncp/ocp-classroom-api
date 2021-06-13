package flusher

import (
	"context"

	"github.com/ozoncp/ocp-classroom-api/internal/models"
	"github.com/ozoncp/ocp-classroom-api/internal/repo"
	"github.com/ozoncp/ocp-classroom-api/internal/utils"
)

// TODO: comment everything here

type Flusher interface {
	Flush(ctx context.Context, classrooms []models.Classroom) []models.Classroom
}

type flusher struct {
	repo      repo.Repo
	chunkSize int
}

func New(repo repo.Repo, chunkSize int) *flusher {
	return &flusher{repo: repo, chunkSize: chunkSize}
}

func (fl *flusher) Flush(ctx context.Context, classrooms []models.Classroom) []models.Classroom {

	chunks, err := utils.SplitSlice(classrooms, fl.chunkSize)

	if err != nil {
		return classrooms
	}

	for i, chunk := range chunks {

		if _, err := fl.repo.MultiAddClassroom(ctx, chunk); err != nil {
			return classrooms[fl.chunkSize*i:]
		}
	}

	return nil
}

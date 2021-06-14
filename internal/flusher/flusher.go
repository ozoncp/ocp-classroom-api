package flusher

import (
	"context"

	opentracing "github.com/opentracing/opentracing-go"

	"github.com/opentracing/opentracing-go/log"

	"github.com/ozoncp/ocp-classroom-api/internal/models"
	"github.com/ozoncp/ocp-classroom-api/internal/repo"
	"github.com/ozoncp/ocp-classroom-api/internal/utils"
)

// Flusher is utility that flushes classrooms to any storage and provides tracing
type Flusher interface {
	Flush(ctx context.Context, span opentracing.Span, classrooms []models.Classroom) []models.Classroom
}

// flusher is implementation of Flusher interface that flushes classrooms to DB and split huge slices to small chunks
type flusher struct {
	repo      repo.Repo
	chunkSize int
}

// New returns flusher instance for flushing classrooms to DB
func New(repo repo.Repo, chunkSize int) *flusher {
	return &flusher{repo: repo, chunkSize: chunkSize}
}

// Flush flushes classrooms to DB by small chunks and provides tracing
func (fl *flusher) Flush(ctx context.Context, span opentracing.Span, classrooms []models.Classroom) []models.Classroom {

	chunks, err := utils.SplitSlice(classrooms, fl.chunkSize)

	if err != nil {
		return classrooms
	}

	for i, chunk := range chunks {

		var childSpan opentracing.Span
		if span != nil {
			childSpan = opentracing.StartSpan("Flush", opentracing.ChildOf(span.Context()))
		}

		_, err := fl.repo.MultiAddClassroom(ctx, chunk)

		if span != nil {

			childSpan.LogFields(
				log.Int("len", len(chunk)),
				log.Bool("sent", err == nil),
			)
			childSpan.Finish()
		}

		if err != nil {
			return classrooms[fl.chunkSize*i:]
		}
	}

	return nil
}

package flusher

import "github.com/ozoncp/ocp-classroom-api/internal/models"

type Flusher interface {
	Flush(tasks []models.Classroom) []models.Classroom
}

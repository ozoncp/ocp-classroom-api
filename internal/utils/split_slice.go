package utils

import (
	"errors"

	"github.com/ozoncp/ocp-classroom-api/internal/models"
)

func SplitSlice(src []models.Classroom, chunkSize int) (dst [][]models.Classroom, err error) {

	if src == nil {
		err = errors.New("src is nil")
		return
	}

	if chunkSize <= 0 {
		err = errors.New("chunkSize <= 0")
		return
	}

	for i := 0; i < len(src)/chunkSize; i++ {

		begin := 0 + chunkSize*i
		end := chunkSize + chunkSize*i

		dst = append(dst, src[begin:end])
	}

	left := len(src) % chunkSize

	if left > 0 {

		dst = append(dst, src[len(src)-left:])
	}

	return
}

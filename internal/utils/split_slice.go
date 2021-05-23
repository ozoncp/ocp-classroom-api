package utils

import "github.com/ozoncp/ocp-classroom-api/internal/models"

func SplitSlice(src []models.Classroom, chunkSize int) (dst [][]models.Classroom) {

	if src == nil || chunkSize <= 0 {
		return
	}

	dst = [][]models.Classroom{}

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

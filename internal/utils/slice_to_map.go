package utils

import (
	"errors"

	"github.com/ozoncp/ocp-classroom-api/internal/models"
)

func SliceToMap(src []models.Classroom) (dst map[uint]models.Classroom, err error) {

	if src == nil {
		err = errors.New("slice is nil")
		return
	}

	dst = make(map[uint]models.Classroom)

	for _, value := range src {

		if _, found := dst[value.ClassroomId()]; found {
			err = errors.New("id is already present")
			return
		}

		dst[value.ClassroomId()] = value
	}

	return
}

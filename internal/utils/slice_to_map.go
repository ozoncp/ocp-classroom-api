package utils

import (
	"errors"

	"github.com/ozoncp/ocp-classroom-api/internal/models"
)

func SliceToMap(src []models.Classroom) (dst map[uint64]models.Classroom, err error) {

	if src == nil {
		err = errors.New("slice is nil")
		return
	}

	if len(src) == 0 {
		return
	}

	dst = make(map[uint64]models.Classroom, len(src))

	for _, value := range src {

		if _, found := dst[value.Id]; found {
			err = errors.New("id is already present")
			return
		}

		dst[value.Id] = value
	}

	return
}

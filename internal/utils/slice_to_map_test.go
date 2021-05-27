package utils

import (
	"errors"
	"reflect"
	"testing"

	"github.com/ozoncp/ocp-classroom-api/internal/models"
)

func TestSliceToMap(t *testing.T) {

	type outData struct {
		dst map[uint]models.Classroom
		err error
	}

	type testCase struct {
		in   []models.Classroom
		want outData
	}

	var testCases = [...]testCase{
		{
			in: nil,

			want: outData{
				nil,
				errors.New("slice is nil"),
			},
		},

		{
			in: []models.Classroom{
				models.New(0, 0, 0),
				models.New(0, 1, 1),
			},

			want: outData{
				map[uint]models.Classroom{
					0: models.New(0, 0, 0),
				},
				errors.New("id is already present"),
			},
		},

		{
			in: []models.Classroom{},

			want: outData{
				nil,
				nil,
			},
		},

		{
			in: []models.Classroom{
				models.New(0, 0, 0),
				models.New(1, 1, 1),
			},

			want: outData{
				map[uint]models.Classroom{
					0: models.New(0, 0, 0),
					1: models.New(1, 1, 1),
				},
				nil,
			},
		},
	}

	defer func() { recover() }()

	for i, testCase := range testCases {

		var got outData
		got.dst, got.err = SliceToMap(testCase.in)

		if !reflect.DeepEqual(testCase.want.dst, got.dst) ||
			(testCase.want.err != nil && got.err == nil) ||
			(testCase.want.err == nil && got.err != nil) {

			t.Errorf("test[%v]: want: %v, got : %v.", i, testCase.want, got)
		}
	}
}
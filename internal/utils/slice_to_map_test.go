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
				models.New(0, nil, nil),
				models.New(0, nil, nil),
			},

			want: outData{
				map[uint]models.Classroom{
					0: models.New(0, nil, nil),
				},
				errors.New("id is already present"),
			},
		},

		{
			in: []models.Classroom{},

			want: outData{
				map[uint]models.Classroom{},
				nil,
			},
		},

		{
			in: []models.Classroom{
				models.New(0, nil, nil),
				models.New(1, []uint{1, 2, 3}, nil),
			},

			want: outData{
				map[uint]models.Classroom{
					0: models.New(0, nil, nil),
					1: models.New(1, []uint{1, 2, 3}, nil),
				},
				nil,
			},
		},
	}

	defer func() { recover() }()

	for _, testCase := range testCases {

		var got outData
		got.dst, got.err = SliceToMap(testCase.in)

		if !reflect.DeepEqual(testCase.want.dst, got.dst) ||
			(testCase.want.err != nil && got.err == nil) ||
			(testCase.want.err == nil && got.err != nil) {

			t.Errorf("want: %v, got : %v.", testCase.want, got)
		}
	}
}

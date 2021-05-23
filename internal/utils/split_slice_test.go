package utils

import (
	"reflect"
	"testing"

	"github.com/ozoncp/ocp-classroom-api/internal/models"
)

func TestSplitSlice(t *testing.T) {

	type inData struct {
		slice     []models.Classroom
		chunkSize int
	}

	type testCase struct {
		in   inData
		want [][]models.Classroom
	}

	var testCases = [...]testCase{

		{
			in: inData{
				slice:     nil,
				chunkSize: 3,
			},

			want: nil,
		},

		{
			in: inData{
				slice:     nil,
				chunkSize: 0,
			},

			want: nil,
		},

		{
			in: inData{
				slice:     []models.Classroom{},
				chunkSize: 3,
			},

			want: [][]models.Classroom{},
		},

		{
			in: inData{
				slice: []models.Classroom{
					models.New(0, nil, nil),
					models.New(1, nil, nil),
					models.New(2, nil, nil),
					models.New(3, nil, nil),
				},
				chunkSize: -1,
			},

			want: nil,
		},

		{
			in: inData{
				slice: []models.Classroom{
					models.New(0, nil, nil),
					models.New(1, nil, nil),
					models.New(2, nil, nil),
					models.New(3, nil, nil),
				},
				chunkSize: 0,
			},

			want: nil,
		},

		{
			in: inData{
				slice: []models.Classroom{
					models.New(0, nil, nil),
					models.New(1, nil, nil),
					models.New(2, nil, nil),
					models.New(3, nil, nil),
					models.New(4, nil, nil),
					models.New(5, nil, nil),
				},
				chunkSize: 3,
			},

			want: [][]models.Classroom{
				{
					models.New(0, nil, nil),
					models.New(1, nil, nil),
					models.New(2, nil, nil),
				},

				{
					models.New(3, nil, nil),
					models.New(4, nil, nil),
					models.New(5, nil, nil),
				},
			},
		},

		{
			in: inData{
				slice: []models.Classroom{
					models.New(0, nil, nil),
					models.New(1, nil, nil),
					models.New(2, nil, nil),
					models.New(3, nil, nil),
					models.New(4, nil, nil),
					models.New(5, nil, nil),
					models.New(6, nil, nil),
				},
				chunkSize: 3,
			},

			want: [][]models.Classroom{
				{
					models.New(0, nil, nil),
					models.New(1, nil, nil),
					models.New(2, nil, nil),
				},

				{
					models.New(3, nil, nil),
					models.New(4, nil, nil),
					models.New(5, nil, nil),
				},

				{
					models.New(6, nil, nil),
				},
			},
		},

		{
			in: inData{
				slice: []models.Classroom{
					models.New(0, []uint{1, 2}, nil),
					models.New(1, nil, nil),
					models.New(2, nil, nil),
				},
				chunkSize: 5,
			},

			want: [][]models.Classroom{
				{
					models.New(0, []uint{1, 2}, nil),
					models.New(1, nil, nil),
					models.New(2, nil, nil),
				},
			},
		},
	}

	for _, testCase := range testCases {

		got := SplitSlice(testCase.in.slice, testCase.in.chunkSize)

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("want: %v, got : %v.", testCase.want, got)
		}
	}
}

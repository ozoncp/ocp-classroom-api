package utils

import (
	"errors"
	"reflect"
	"testing"

	"github.com/ozoncp/ocp-classroom-api/internal/models"
)

func TestSplitSlice(t *testing.T) {

	type inData struct {
		src       []models.Classroom
		chunkSize int
	}

	type outData struct {
		dst [][]models.Classroom
		err error
	}

	type testCase struct {
		in   inData
		want outData
	}

	var testCases = [...]testCase{

		{
			in: inData{
				src:       nil,
				chunkSize: 3,
			},

			want: outData{
				dst: nil,
				err: errors.New("src is nil"),
			},
		},

		{
			in: inData{
				src:       []models.Classroom{},
				chunkSize: 3,
			},

			want: outData{
				dst: nil,
				err: nil,
			},
		},

		{
			in: inData{
				src: []models.Classroom{
					models.New(0, nil, nil),
					models.New(1, nil, nil),
					models.New(2, nil, nil),
					models.New(3, nil, nil),
				},
				chunkSize: -1,
			},

			want: outData{
				dst: nil,
				err: errors.New("chunkSize <= 0"),
			},
		},

		{
			in: inData{
				src: []models.Classroom{
					models.New(0, nil, nil),
					models.New(1, nil, nil),
					models.New(2, nil, nil),
					models.New(3, nil, nil),
					models.New(4, nil, nil),
					models.New(5, nil, nil),
				},
				chunkSize: 3,
			},

			want: outData{
				dst: [][]models.Classroom{
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
				err: nil,
			},
		},

		{
			in: inData{
				src: []models.Classroom{
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

			want: outData{
				dst: [][]models.Classroom{
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
				err: nil,
			},
		},

		{
			in: inData{
				src: []models.Classroom{
					models.New(0, []uint{1, 2}, nil),
					models.New(1, nil, nil),
					models.New(2, nil, nil),
				},
				chunkSize: 5,
			},

			want: outData{
				dst: [][]models.Classroom{
					{
						models.New(0, []uint{1, 2}, nil),
						models.New(1, nil, nil),
						models.New(2, nil, nil),
					},
				},
				err: nil,
			},
		},
	}

	for i, testCase := range testCases {

		var got outData
		got.dst, got.err = SplitSlice(testCase.in.src, testCase.in.chunkSize)

		if !reflect.DeepEqual(testCase.want.dst, got.dst) ||
			(testCase.want.err != nil && got.err == nil) ||
			(testCase.want.err == nil && got.err != nil) {

			t.Errorf("test[%v]: want: %v, got : %v.", i, testCase.want, got)
		}
	}
}

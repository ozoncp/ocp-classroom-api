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
					{Id: 0, TenantId: 0, CalendarId: 0},
					{Id: 1, TenantId: 1, CalendarId: 1},
					{Id: 2, TenantId: 2, CalendarId: 2},
					{Id: 3, TenantId: 3, CalendarId: 3},
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
					{Id: 0, TenantId: 0, CalendarId: 0},
					{Id: 1, TenantId: 1, CalendarId: 1},
					{Id: 2, TenantId: 2, CalendarId: 2},
					{Id: 3, TenantId: 3, CalendarId: 3},
					{Id: 4, TenantId: 4, CalendarId: 4},
					{Id: 5, TenantId: 5, CalendarId: 5},
				},
				chunkSize: 3,
			},

			want: outData{
				dst: [][]models.Classroom{
					{
						{Id: 0, TenantId: 0, CalendarId: 0},
						{Id: 1, TenantId: 1, CalendarId: 1},
						{Id: 2, TenantId: 2, CalendarId: 2},
					},

					{
						{Id: 3, TenantId: 3, CalendarId: 3},
						{Id: 4, TenantId: 4, CalendarId: 4},
						{Id: 5, TenantId: 5, CalendarId: 5},
					},
				},
				err: nil,
			},
		},

		{
			in: inData{
				src: []models.Classroom{
					{Id: 0, TenantId: 0, CalendarId: 0},
					{Id: 1, TenantId: 1, CalendarId: 1},
					{Id: 2, TenantId: 2, CalendarId: 2},
					{Id: 3, TenantId: 3, CalendarId: 3},
					{Id: 4, TenantId: 4, CalendarId: 4},
					{Id: 5, TenantId: 5, CalendarId: 5},
					{Id: 6, TenantId: 6, CalendarId: 6},
				},
				chunkSize: 3,
			},

			want: outData{
				dst: [][]models.Classroom{
					{
						{Id: 0, TenantId: 0, CalendarId: 0},
						{Id: 1, TenantId: 1, CalendarId: 1},
						{Id: 2, TenantId: 2, CalendarId: 2},
					},

					{
						{Id: 3, TenantId: 3, CalendarId: 3},
						{Id: 4, TenantId: 4, CalendarId: 4},
						{Id: 5, TenantId: 5, CalendarId: 5},
					},

					{
						{Id: 6, TenantId: 6, CalendarId: 6},
					},
				},
				err: nil,
			},
		},

		{
			in: inData{
				src: []models.Classroom{
					{Id: 0, TenantId: 0, CalendarId: 0},
					{Id: 1, TenantId: 1, CalendarId: 1},
					{Id: 2, TenantId: 2, CalendarId: 2},
				},
				chunkSize: 5,
			},

			want: outData{
				dst: [][]models.Classroom{
					{
						{Id: 0, TenantId: 0, CalendarId: 0},
						{Id: 1, TenantId: 1, CalendarId: 1},
						{Id: 2, TenantId: 2, CalendarId: 2},
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

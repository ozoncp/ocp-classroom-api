package utils

import (
	"errors"
	"reflect"
	"testing"

	"github.com/ozoncp/ocp-classroom-api/internal/models"
)

func TestSliceToMap(t *testing.T) {

	type outData struct {
		dst map[uint64]models.Classroom
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
				{Id: 0, TenantId: 0, CalendarId: 0},
				{Id: 0, TenantId: 1, CalendarId: 1},
			},

			want: outData{
				map[uint64]models.Classroom{
					0: {Id: 0, TenantId: 0, CalendarId: 0},
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
				{Id: 0, TenantId: 0, CalendarId: 0},
				{Id: 1, TenantId: 1, CalendarId: 1},
			},

			want: outData{
				map[uint64]models.Classroom{
					0: {Id: 0, TenantId: 0, CalendarId: 0},
					1: {Id: 1, TenantId: 1, CalendarId: 1},
				},
				nil,
			},
		},
	}

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

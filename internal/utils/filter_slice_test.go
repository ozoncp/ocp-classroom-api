package utils

import (
	"reflect"
	"testing"
)

func TestFilterSlice(t *testing.T) {

	type inData struct {
		src  []int
		keys []int
	}

	type testCase struct {
		in   inData
		want []int
	}

	var testCases = [...]testCase{
		{
			in: inData{
				src:  nil,
				keys: nil,
			},

			want: nil,
		},

		{
			in: inData{
				src:  []int{},
				keys: nil,
			},

			want: nil,
		},

		{
			in: inData{
				src:  nil,
				keys: []int{},
			},

			want: nil,
		},

		{
			in: inData{
				src:  []int{},
				keys: []int{},
			},

			want: []int{},
		},

		{
			in: inData{
				src:  []int{1, 2, 3, 4, 5, 6},
				keys: []int{},
			},

			want: []int{1, 2, 3, 4, 5, 6},
		},

		{
			in: inData{
				src:  []int{1, 2, 3, 4, 5, 6},
				keys: []int{1, 2, 3},
			},

			want: []int{4, 5, 6},
		},

		{
			in: inData{
				src:  []int{4, 5, 6},
				keys: []int{1, 2, 3},
			},

			want: []int{4, 5, 6},
		},

		{
			in: inData{
				src:  []int{1, 2, 3},
				keys: []int{1, 2, 3},
			},

			want: []int{},
		},
	}

	for _, testCase := range testCases {
		got := FilterSlice(testCase.in.src, testCase.in.keys)

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("want: %v, got : %v.", testCase.want, got)
		}
	}
}

package utils

import (
	"errors"
	"reflect"
	"testing"
)

func TestFilterSlice(t *testing.T) {

	type inData struct {
		src  []int
		keys []int
	}

	type outData struct {
		dst []int
		err error
	}

	type testCase struct {
		in   inData
		want outData
	}

	var testCases = [...]testCase{
		{
			in: inData{
				src:  nil,
				keys: nil,
			},

			want: outData{
				nil,
				errors.New("src is nil"),
			},
		},

		{
			in: inData{
				src:  []int{},
				keys: nil,
			},

			want: outData{
				nil,
				errors.New("keys is nil"),
			},
		},

		{
			in: inData{
				src:  []int{},
				keys: []int{},
			},

			want: outData{
				nil,
				nil,
			},
		},

		{
			in: inData{
				src:  []int{1, 2, 3, 4, 5, 6},
				keys: []int{},
			},

			want: outData{
				[]int{1, 2, 3, 4, 5, 6},
				nil,
			},
		},

		{
			in: inData{
				src:  []int{1, 2, 3, 4, 5, 6},
				keys: []int{1, 2, 3},
			},

			want: outData{
				[]int{4, 5, 6},
				nil,
			},
		},

		{
			in: inData{
				src:  []int{4, 5, 6},
				keys: []int{1, 2, 3},
			},

			want: outData{
				[]int{4, 5, 6},
				nil,
			},
		},

		{
			in: inData{
				src:  []int{1, 2, 3},
				keys: []int{1, 2, 3},
			},

			want: outData{
				nil,
				nil,
			},
		},
	}

	for i, testCase := range testCases {

		var got outData
		got.dst, got.err = FilterSlice(testCase.in.src, testCase.in.keys)

		if !reflect.DeepEqual(testCase.want.dst, got.dst) ||
			(testCase.want.err != nil && got.err == nil) ||
			(testCase.want.err == nil && got.err != nil) {

			t.Errorf("test[%v]: want: %v, got : %v.", i, testCase.want, got)
		}
	}
}

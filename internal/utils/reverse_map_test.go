package utils

import (
	"errors"
	"reflect"
	"testing"
)

func TestReverseMap(t *testing.T) {

	type outData struct {
		dst map[string]int
		err error
	}

	type testCase struct {
		in   map[int]string
		want outData
	}

	var testCases = [...]testCase{
		{
			in: nil,

			want: outData{
				nil,
				errors.New("src is nil"),
			},
		},

		{
			in: map[int]string{},

			want: outData{
				nil,
				nil,
			},
		},

		{
			in: map[int]string{
				1: "one",
				2: "two",
				3: "three",
			},

			want: outData{
				map[string]int{
					"one":   1,
					"two":   2,
					"three": 3,
				},
				nil,
			},
		},

		{
			in: map[int]string{
				1: "wait for error",
				2: "wait for error",
				3: "wait for error",
			},

			want: outData{
				map[string]int{
					"wait for error": 1,
				},
				errors.New("key is already present"),
			},
		},
	}

	for i, testCase := range testCases {

		var got outData
		got.dst, got.err = ReverseMap(testCase.in)

		if !reflect.DeepEqual(testCase.want.dst, got.dst) ||
			(testCase.want.err != nil && got.err == nil) ||
			(testCase.want.err == nil && got.err != nil) {

			t.Errorf("test[%v]: want: %v, got : %v.", i, testCase.want, got)
		}
	}
}

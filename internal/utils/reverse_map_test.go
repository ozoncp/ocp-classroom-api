package utils

import (
	"reflect"
	"testing"
)

func TestReverseMap(t *testing.T) {

	type testCase struct {
		in   map[int]string
		want map[string]int
	}

	var testCases = [...]testCase{
		{
			in: nil,

			want: nil,
		},

		{
			in: map[int]string{},

			want: map[string]int{},
		},

		{
			in: map[int]string{
				1: "one",
				2: "two",
				3: "three",
			},

			want: map[string]int{
				"one":   1,
				"two":   2,
				"three": 3,
			},
		},

		// panic test
		{
			in: map[int]string{
				1: "wait for panic",
				2: "wait for panic",
				3: "wait for panic",
			},

			want: map[string]int{
				"wait for panic": 1,
			},
		},
	}

	defer func() { recover() }()

	for _, testCase := range testCases {
		got := ReverseMap(testCase.in)

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("want: %v, got : %v.", testCase.want, got)
		}
	}
}

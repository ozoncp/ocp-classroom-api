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

		// Здесь не уверен!
		// Я так понимаю, что результат же может быть разным из-за unordered свойства map? :C
		// Помогите 🐱
		{
			in: map[int]string{
				1: "Im'confused",
				2: "Im'confused",
				3: "Im'confused",
			},

			want: map[string]int{
				"Im'confused": 3,
			},
		},
	}

	for _, testCase := range testCases {
		got := ReverseMap(testCase.in)

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("want: %v, got : %v.", testCase.want, got)
		}
	}
}

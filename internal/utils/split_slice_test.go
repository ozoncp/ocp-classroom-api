package utils

import (
	"reflect"
	"testing"
)

func TestSplitSliceToChunks(t *testing.T) {

	type inData struct {
		slice     []int
		chunkSize int
	}

	type testCase struct {
		in   inData
		want [][]int
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
				slice:     []int{},
				chunkSize: 3,
			},

			want: [][]int{},
		},

		{
			in: inData{
				slice:     []int{0, 1, 2, 4},
				chunkSize: -1,
			},

			want: nil,
		},

		{
			in: inData{
				slice:     []int{0, 1, 2, 4},
				chunkSize: 0,
			},

			want: nil,
		},

		{
			in: inData{
				slice:     []int{0, 1, 2, 3, 4, 5},
				chunkSize: 3,
			},

			want: [][]int{
				{0, 1, 2},
				{3, 4, 5},
			},
		},

		{
			in: inData{
				slice:     []int{0, 1, 2, 3, 4, 5, 6, 7},
				chunkSize: 3,
			},

			want: [][]int{
				{0, 1, 2},
				{3, 4, 5},
				{6, 7},
			},
		},

		{
			in: inData{
				slice:     []int{0, 1, 2, 3, 4, 5, 6, 7},
				chunkSize: 10,
			},

			want: [][]int{
				{0, 1, 2, 3, 4, 5, 6, 7},
			},
		},
	}

	for _, testCase := range testCases {

		got := SplitSliceToChunks(testCase.in.slice, testCase.in.chunkSize)

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("want: %v, got : %v.", testCase.want, got)
		}
	}
}

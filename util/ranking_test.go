package util

import (
	"reflect"
	"testing"
)

func TestRankInt(t *testing.T) {
	testCases := []struct {
		name     string
		scores   []int
		reversed bool
		want     []int
	}{
		{
			name:     "example test case",
			scores:   []int{9, 8, 10, 7, 8, 9},
			reversed: false,
			want:     []int{2, 4, 1, 6, 4, 2},
		},
		{
			name:     "ascending test case",
			scores:   []int{1, 2, 3, 4, 5, 6},
			reversed: false,
			want:     []int{6, 5, 4, 3, 2, 1},
		},
		{
			name:     "descending test case",
			scores:   []int{6, 5, 4, 3, 2, 1},
			reversed: false,
			want:     []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "same value test case",
			scores:   []int{5, 5, 5, 5, 5, 5},
			reversed: false,
			want:     []int{1, 1, 1, 1, 1, 1},
		},
		{
			name:     "example reversed test case",
			scores:   []int{9, 8, 10, 7, 8, 9},
			reversed: true,
			want:     []int{4, 2, 6, 1, 2, 4},
		},
		{
			name:     "ascending reversed test case",
			scores:   []int{1, 2, 3, 4, 5, 6},
			reversed: true,
			want:     []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "descending reversed test case",
			scores:   []int{6, 5, 4, 3, 2, 1},
			reversed: true,
			want:     []int{6, 5, 4, 3, 2, 1},
		},
		{
			name:     "same value reversed test case",
			scores:   []int{5, 5, 5, 5, 5, 5},
			reversed: true,
			want:     []int{1, 1, 1, 1, 1, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Rank(tc.scores, tc.reversed)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Rank(%v) = %v, want %v", tc.scores, got, tc.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "even length input",
			input: []int{1, 2, 3, 4},
			want:  []int{4, 3, 2, 1},
		},
		{
			name:  "odd length input",
			input: []int{1, 2, 3},
			want:  []int{3, 2, 1},
		},
		{
			name:  "empty input",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "single element input",
			input: []int{5},
			want:  []int{5},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := make([]int, len(tc.input))
			copy(input, tc.input)
			reverse(input)
			if !reflect.DeepEqual(input, tc.want) {
				t.Errorf("reverse(%v) = %v, want %v", tc.input, input, tc.want)
			}
		})
	}
}

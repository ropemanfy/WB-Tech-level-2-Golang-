package main

import (
	"testing"
)

type TestData struct {
	str   []string
	flags Flags
	want  []string
}

func TestManSort(t *testing.T) {

	tests := []TestData{
		{
			str: []string{"хлеб 1 черепаха", "колбаса 7 лама", "гвозди", "колбаса 7 лама", "хлеб 4 коала"},
			flags: Flags{
				column:  0,
				byNum:   false,
				reverse: false,
				uniq:    false,
			},
			want: []string{"гвозди", "колбаса 7 лама", "колбаса 7 лама", "хлеб 1 черепаха", "хлеб 4 коала"},
		},
		{
			str: []string{"хлеб 1 черепаха", "колбаса 7 лама", "гвозди", "колбаса 7 лама", "хлеб 4 коала"},
			flags: Flags{
				column:  3,
				byNum:   false,
				reverse: false,
				uniq:    false,
			},
			want: []string{"гвозди", "хлеб 4 коала", "колбаса 7 лама", "колбаса 7 лама", "хлеб 1 черепаха"},
		},
	}

	for num, test := range tests {
		result := manSort(test.str, test.flags)
		if len(result) != len(test.want) {
			t.Errorf("Test %v: result != want", (num + 1))
			return
		}
		for i, v := range result {
			if test.want[i] != result[i] {
				t.Errorf("Test %v: got = %s, want %s", (num + 1), v, test.want[i])
			}
		}
	}
}

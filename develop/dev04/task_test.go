package main

import (
	"slices"
	"testing"
)

type TestData struct {
	str  []string
	want map[string][]string
}

func TestAnagramm(t *testing.T) {
	tests := []TestData{
		{
			str: []string{
				"пятка", "пяТак", "тяпКа",
				"листок", "слиток", "столик", "апельсин",
				"спаниель", "торг", "грот",
			},
			want: map[string][]string{
				"пятка":    {"пятка", "пятак", "тяпка"},
				"листок":   {"листок", "слиток", "столик"},
				"апельсин": {"апельсин", "спаниель"},
				"торг":     {"торг", "грот"},
			},
		},
		{
			str: []string{
				"пятка", "ы", "тяпКа",
				"", "", "столик", "торг", "апельсин",
				"спаниель", "торг", "грот",
			},
			want: map[string][]string{
				"пятка":    {"пятка", "тяпка"},
				"апельсин": {"апельсин", "спаниель"},
				"торг":     {"торг", "грот"},
			},
		},
	}

	for num, test := range tests {
		result := Anagramm(test.str)
		if len(result) != len(test.want) {
			t.Errorf("Test %v: result != want", (num + 1))
			return
		}
		for k, v := range result {
			if !slices.Equal(v, test.want[k]) {
				t.Errorf("Test %v: got = %s, want %s", (num + 1), v, test.want[k])
			}
		}
	}
}

package main

import "testing"

type TestData struct {
	fileName string
	str      string
	flags    Flags
	want     []string
}

func TestManSort(t *testing.T) {

	tests := []TestData{
		{
			fileName: "file.txt",
			str:      "весна",
			flags:    Flags{},
			want:     []string{"весна цветы"},
		},
		{
			fileName: "file.txt",
			str:      "весна",
			flags:    Flags{after: 2},
			want:     []string{"весна цветы", "Лето луг", "зима лёд"},
		},
		{
			fileName: "file.txt",
			str:      "весна",
			flags:    Flags{before: 1},
			want:     []string{"осень дождь", "весна цветы"},
		},
		{
			fileName: "file.txt",
			str:      "весна",
			flags:    Flags{context: 2},
			want:     []string{"осень дождь", "весна цветы", "Лето луг"},
		},
		{
			fileName: "file.txt",
			str:      "луг",
			flags:    Flags{count: true},
			want:     []string{"лето луг", "Лето луг", "количество совпадений: 2"},
		},
		{
			fileName: "file.txt",
			str:      "лето",
			flags:    Flags{ignoreCase: true},
			want:     []string{"лето луг", "Лето луг"},
		},
		{
			fileName: "file.txt",
			str:      "весна",
			flags:    Flags{invert: true},
			want:     []string{"лето луг", "осень дождь", "осень дождь", "Лето луг", "зима лёд", "Осень дождь"},
		},
		{
			fileName: "file.txt",
			str:      "Лето луг",
			flags:    Flags{fixed: true},
			want:     []string{"Лето луг"},
		},
		{
			fileName: "file.txt",
			str:      "луг",
			flags:    Flags{lineNum: true},
			want:     []string{"1: лето луг", "5: Лето луг"},
		},
	}

	for num, test := range tests {
		result := manGrep(test.fileName, test.str, test.flags)
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

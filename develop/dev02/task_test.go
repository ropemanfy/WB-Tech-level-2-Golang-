package main

import "testing"

type TestData struct {
	str  string
	want string
}

func TestUnpack(t *testing.T) {
	tests := []TestData{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"45", ""},
		{"", ""},
		{"qwe\\4\\5", "qwe45"},
		{"qwe\\45", "qwe44444"},
		{"qwe\\\\5", "qwe\\\\\\\\\\"},
	}

	for num, test := range tests {
		result, _ := Unpack(test.str)
		if result != test.want {
			t.Errorf("Test %v: got = %s, want %s", (num + 1), result, test.want)
		}
	}
}

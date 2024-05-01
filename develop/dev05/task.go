package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Flags struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
}

func main() {
	flags := Flags{}
	flag.IntVar(&flags.after, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&flags.before, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&flags.context, "C", 0, "печатать ±N строк вокруг совпадения")
	flag.BoolVar(&flags.count, "c", false, "количество строк")
	flag.BoolVar(&flags.ignoreCase, "i", false, "игнорировать регистр")
	flag.BoolVar(&flags.invert, "v", false, "вместо совпадения, исключать")
	flag.BoolVar(&flags.fixed, "F", false, "точное совпадение со строкой, не паттерн")
	flag.BoolVar(&flags.lineNum, "n", false, "печатать номер строки")
	flag.Parse()

	file := flag.Arg(0)
	substr := flag.Arg(1)

	result := manGrep(file, substr, flags)
	for _, v := range result {
		fmt.Println(v)
	}
}

func manGrep(fileName, substr string, flags Flags) (result []string) {
	text := getStrings(fileName)

	matches, count := match(text, substr, flags)

	if flags.after > 0 || flags.before > 0 || flags.context > 0 {
		matches = addOffset(matches, len(text), flags)
	}

	if flags.invert {
		matches = invert(matches, text)
	}

	for i, v := range text {
		if slices.Contains(matches, i) {
			switch {
			case flags.lineNum:
				v = fmt.Sprintf("%v: ", (i+1)) + v
				result = append(result, v)
			default:
				result = append(result, v)
			}
		}
	}

	if flags.count {
		result = append(result, fmt.Sprintf("количество совпадений: %v", count))
	}
	return
}

func getStrings(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("open file: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("read file: %v", err)
	}

	text := strings.Split(string(data), "\r\n")
	return text
}

func match(text []string, substr string, flags Flags) (result []int, count int) {
	for i, str := range text {
		if flags.ignoreCase {
			str = strings.ToLower(str)
			substr = strings.ToLower(substr)
		}
		switch {
		case flags.fixed:
			if str == substr {
				result = append(result, i)
			}
		default:
			if strings.Contains(str, substr) {
				result = append(result, i)
			}
		}
	}
	count = len(result)
	return
}

func addOffset(matches []int, border int, flags Flags) (result []int) {
	after, before := flags.after, flags.before
	if flags.context > 0 {
		after = flags.context / 2
		before = (flags.context / 2) + (flags.context % 2)
	}

	for _, v := range matches {
		for i := (v - before); i < v; i++ {
			if i >= 0 {
				result = append(result, i)
			}
		}
		result = append(result, v)
		for i := (v + after); i > v; i-- {
			if i < border {
				result = append(result, i)
			}
		}
	}
	return
}

func invert(matches []int, text []string) (result []int) {
	for i := range text {
		if !slices.Contains(matches, i) {
			result = append(result, i)
		}
	}
	return
}

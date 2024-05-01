package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Flags struct {
	fields    int
	delimiter string
	separated bool
}

func main() {
	flags := Flags{}
	flag.IntVar(&flags.fields, "f", 0, "выбрать поля (колонки)")
	flag.StringVar(&flags.delimiter, "d", "\t", "использовать другой разделитель")
	flag.BoolVar(&flags.separated, "s", false, "только строки с разделителем")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(`Введите текст. Для старта утилиты введите "cut"`)

	var text []string
	isRun := true
	for isRun {
		scanner.Scan()
		str := scanner.Text()
		switch str {
		case "cut":
			isRun = false
		default:
			text = append(text, str)
		}
	}

	for _, v := range text {
		fmt.Print(manCut(flags, v))
	}
}

func manCut(flags Flags, str string) (result string) {
	if flags.separated && !strings.Contains(str, flags.delimiter) {
		return
	}

	words := strings.Split(str, flags.delimiter)
	if flags.fields <= len(words) {
		var sb strings.Builder
		field := getField(flags)
		sb.WriteString(words[field] + "\n")
		result = sb.String()
		return
	}
	return
}

func getField(flags Flags) (column int) {
	if flags.fields > 0 {
		column = flags.fields - 1
	}
	return
}

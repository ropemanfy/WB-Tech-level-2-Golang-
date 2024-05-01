package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	s := "a4bc2d5e"
	fmt.Println(Unpack(s))
}

func Unpack(s string) (string, error) {
	if _, err := strconv.Atoi(s); err == nil {
		err = fmt.Errorf("invalid string")
		return "", err
	}

	const command = "\\"
	var isCommand = false
	var sb strings.Builder
	var char rune

	for _, v := range s {
		count, err := strconv.Atoi(string(v))
		if err != nil {
			switch {
			case string(v) == command && !isCommand:
				isCommand = true
				continue
			case string(v) == command && isCommand:
				isCommand = false
			}
			char = v
			sb.WriteRune(char)
		}
		if isCommand {
			char = v
			count = 0
			sb.WriteRune(char)
			isCommand = false
		}
		for i := 0; i < count-1; i++ {
			sb.WriteRune(char)
		}
	}
	return sb.String(), nil
}

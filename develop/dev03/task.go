package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Flags struct {
	column  int
	byNum   bool
	reverse bool
	uniq    bool
}

func main() {
	flags := Flags{}
	flag.IntVar(&flags.column, "k", 0, "указание колонки для сортировки")
	flag.BoolVar(&flags.byNum, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&flags.reverse, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&flags.uniq, "u", false, "не выводить повторяющиеся строки")
	flag.Parse()

	fileName := flag.Arg(0)
	strs := getStrings(fileName)
	strs = manSort(strs, flags)

	err := createFile(strs)
	if err != nil {
		log.Fatalf("create file: %v", err)
	}
}

func createFile(strs []string) error {
	oldName := flag.Arg(0)
	newName := fmt.Sprintf("sorted-%s", oldName)

	f, err := os.Create(newName)
	if err != nil {
		return err
	}

	_, err = f.WriteString(strings.Join(strs, "\n"))
	if err != nil {
		return err
	}

	return nil
}

func manSort(strs []string, flags Flags) (result []string) {
	column := getColumn(flags)

	switch {
	case flags.byNum:
		sort.Slice(strs, func(i, j int) bool {
			numI := getNum(strs[i])
			numJ := getNum(strs[j])
			return func() bool {
				if flags.reverse {
					return numI > numJ
				}
				return numI < numJ
			}()
		})
	default:
		sort.Slice(strs, func(i, j int) bool {
			fieldI := strings.Fields(strs[i])
			fieldJ := strings.Fields(strs[j])

			if column >= len(fieldI) {
				return !(flags.reverse)
			}
			if column >= len(fieldJ) {
				return flags.reverse
			}

			return func() bool {
				if flags.reverse {
					return fieldI[column] > fieldJ[column]
				}
				return fieldI[column] < fieldJ[column]
			}()
		})
	}
	if flags.uniq {
		result = uniq(strs)
	} else {
		result = strs
	}
	return
}

func uniq(str []string) []string {
	result := make([]string, 0, len(str))
	uniqSet := make(map[string]bool)
	for _, v := range str {
		if !uniqSet[v] {
			uniqSet[v] = true
			result = append(result, v)
		}
	}
	return result
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

	strs := strings.Split(string(data), "\r\n")
	return strs
}

func getColumn(flags Flags) (column int) {
	if flags.column > 0 {
		column = flags.column - 1
	}
	return
}

func getNum(s string) int {
	str := strings.Fields(s)
	var num int
	var err error
	for _, v := range str {
		num, err = strconv.Atoi(v)
		if err == nil {
			break
		}
	}
	return num
}

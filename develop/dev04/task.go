package main

import (
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func Anagramm(arr []string) map[string][]string {
	anSets := make(map[string][]string)
	uniqSet := make(map[string]struct{})

	for _, v := range arr {
		v = strings.ToLower(v)
		anagramm := sortLetters(v)
		if _, ok := uniqSet[v]; !ok {
			anSets[anagramm] = append(anSets[anagramm], v)
			uniqSet[v] = struct{}{}
		}
	}

	result := make(map[string][]string)
	for _, v := range anSets {
		if len(v) > 1 {
			result[v[0]] = v
		}
	}

	return result
}

func sortLetters(s string) string {
	split := strings.Split(s, "")
	sort.Strings(split)
	result := strings.Join(split, "")
	return result
}

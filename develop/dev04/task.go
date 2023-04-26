package main

import (
	"fmt"
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

func FindAnagramSets(words *[]string) map[string][]string {
	anagramSets := make(map[string][]string)

	tempMap := make(map[string][]string)
	for _, word := range *words {
		loweredWord := strings.ToLower(word)
		sortedWord := sortString(loweredWord)
		tempMap[sortedWord] = append(tempMap[sortedWord], loweredWord)
	}

	for _, wordSet := range tempMap {
		if len(wordSet) > 1 {
			key := wordSet[0]
			sort.Strings(wordSet)
			anagramSets[key] = wordSet
		}
	}

	return anagramSets
}

func sortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func main() {
	input := []string{"пятка", "пятак", "тяпка", "Слиток", "листок", "столик", "мышь"}
	anagramSets := FindAnagramSets(&input)
	fmt.Printf("%#v\n", anagramSets)
}

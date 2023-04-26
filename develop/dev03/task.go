package main

import (
	"bufio"
	"flag"
	"fmt"
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

func getStringFromColumn(s string, c uint) string {
	fields := strings.Fields(s)
	if len(fields)-1 >= int(c) {
		return fields[c]
	} else {
		return ""
	}
}

func ReadLinesFromFile(file *os.File, uniq bool) []string {
	var fileString string
	uniqnessMap := make(map[string]struct{})

	fileData := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileString = scanner.Text()
		if uniq {
			if _, ok := uniqnessMap[fileString]; !ok {
				uniqnessMap[fileString] = struct{}{}
			} else {
				continue
			}
		}
		fileData = append(fileData, fileString)
	}

	return fileData
}

func SortLines(lines []string, numeric bool, column uint, reverse bool) {
	if numeric {
		sort.Slice(lines, func(i int, j int) bool {
			colI := getStringFromColumn(lines[i], column)
			colJ := getStringFromColumn(lines[j], column)

			in, _ := strconv.Atoi(colI)
			jn, _ := strconv.Atoi(colJ)
			return in < jn
		})
	} else {
		sort.Slice(lines, func(i int, j int) bool {
			colI := getStringFromColumn(lines[i], column)
			colJ := getStringFromColumn(lines[j], column)

			return strings.ToLower(colI) < strings.ToLower(colJ)
		})
	}

	if reverse {
		for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
			lines[i], lines[j] = lines[j], lines[i]
		}
	}
}

func main() {
	// Флаги
	column := flag.Uint("k", 0, "sort via a key")
	numeric := flag.Bool("n", false, "compare according to string numerical value")
	reverse := flag.Bool("r", false, "reverse the result of comparisons")
	uniq := flag.Bool("u", false, "output only the first of an equal run")

	flag.Parse()

	// Читаем файл
	fileData := ReadLinesFromFile(os.Stdin, *uniq)

	// Сортируем
	SortLines(fileData, *numeric, *column, *reverse)

	// Выводим
	for _, s := range fileData {
		fmt.Println(s)
	}
}

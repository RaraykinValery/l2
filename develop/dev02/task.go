package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
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

func Unpack(l string) (string, error) {
	var sbuilder strings.Builder

	rl := []rune(l)

	if len(rl) == 0 {
		return "", nil
	}

	if unicode.IsDigit(rl[0]) {
		return "", errors.New("Digit can't be on the first place in the string")
	}

	for i := 0; i < len(rl); i++ {
		if unicode.IsDigit(rl[i]) {
			num := make([]rune, 0)

			for ; i < len(rl) && unicode.IsDigit(rl[i]); i++ {
				num = append(num, rl[i])
			}
			i--

			n, err := strconv.Atoi(string(num))
			if err != nil {
				return "", errors.New("Can't convert runes to digit")
			}

			for j := 0; j < n-1; j++ {
				sbuilder.WriteRune(rl[i-1])
			}
		} else {
			sbuilder.WriteRune(rl[i])
		}
	}

	return sbuilder.String(), nil
}

func main() {
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		line := input.Text()
		unpackedLine, err := Unpack(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't unpack the line %q: %s\n", line, err)
			os.Exit(1)
		}
		fmt.Printf("origin: %s, unpacked: %s", line, unpackedLine)
	}
}

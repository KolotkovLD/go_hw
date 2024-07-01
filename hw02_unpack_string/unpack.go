package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

// тесты сбиваются
func Unpack(input string) (string, error) {
	var prvChar string
	// Разбиваем входную строку на массив символов
	characters := strings.Split(input, "")

	// Создаем слайс для хранения результата
	result := make([]string, 0)

	// Производим повторение символов
	for ind, char := range characters {
		intChar, err := strconv.Atoi(char)
		if err != nil {
			result = append(result, char)
			prvChar = char
			continue
		}
		// проверка на первую цифру или проверка на число больше 9
		if ind == 0 || prvChar == "" {
			return "", ErrInvalidString
		}

		repeatVal := intChar - 1
		if repeatVal < 0 {
			// проверка на отрицательное число повторений с удалением символа
			result = result[:len(result)-1]
		} else {
			result = append(result, strings.Repeat(prvChar, repeatVal))
			prvChar = ""
		}
	}

	// Объединяем результаты в одну строку
	return strings.Join(result, ""), nil
}

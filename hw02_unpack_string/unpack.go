package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var prvChar string = ""
	// Разбиваем входную строку на массив символов
	characters := strings.Split(input, "")

	// Создаем слайс для хранения результата
	result := make([]string, 0)

	// Производим повторение символов
	for _, char := range characters {
		intChar, err := strconv.Atoi(char)
		if err == nil {
			repeatVal := intChar - 1
			result = append(result, strings.Repeat(prvChar, repeatVal))
		} else {
			result = append(result, char)
			prvChar = char
		}
	}

	// Объединяем результаты в одну строку
	return strings.Join(result, ""), nil
}

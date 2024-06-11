package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	// Разбиваем текст на слова, разделяемые пробелами
	words := strings.Fields(text)

	// Словарь для подсчета частоты каждого слова
	wordCounts := make(map[string]int)
	for _, word := range words {
		wordCounts[strings.ToLower(word)]++
	}

	// Структура для сортировки
	type wordCount struct {
		Word   string
		Count  int
		Length int
	}

	// Создаем слайс для сортировки
	sortedWords := make([]wordCount, 0, len(wordCounts))
	for word, count := range wordCounts {
		sortedWords = append(sortedWords, wordCount{Word: word, Count: count, Length: len(word)})
	}

	// Сортировка слайса по убыванию частоты и лексикографически
	sort.Slice(sortedWords, func(i, j int) bool {
		if sortedWords[i].Count != sortedWords[j].Count {
			return sortedWords[i].Count > sortedWords[j].Count
		}
		return sortedWords[i].Length < sortedWords[j].Length
	})

	// Возвращаем первые 10 слов
	topWords := make([]string, 0, len(wordCounts))
	for i := 0; i < 10 && i < len(sortedWords); i++ {
		//topWords[sortedWords[i].Word] = sortedWords[i].Count
		topWords = append(topWords, sortedWords[i].Word)
	}

	return topWords
}

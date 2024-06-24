package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

// Структура для сортировки

type wordCount struct {
	Word string

	Count int

	Length int
}

func GetMaxWord(listWords []wordCount) wordCount {
	var minWord = listWords[0]

	for i := 0; i < len(listWords); i++ {
		if listWords[i].Word < minWord.Word {
			minWord = listWords[i]
		}
	}

	return minWord
}

func getSameFrequency(words []wordCount, ind int) []wordCount {
	sameFreqWords := make([]wordCount, 0, len(words))

	for i := 0; i < len(words); i++ {
		if ind == words[i].Count {
			sameFreqWords = append(sameFreqWords, words[i])
		}
	}

	return sameFreqWords
}

func deleteWord(words []wordCount, word string) []wordCount {
	rtWords := make([]wordCount, 0, len(words))

	for i := 0; i < len(words); i++ {
		if word != words[i].Word {
			rtWords = append(rtWords, words[i])
		}
	}

	return rtWords
}

func sortSameFrequencyWord(words []wordCount, frqnc int, ind int) ([]string, int) {
	sameFreqWords := getSameFrequency(words, frqnc)

	sortedWords := make([]string, 0, len(words))

	lenSFWs := len(sameFreqWords)

	for i := 0; (i < lenSFWs) && (ind+i < 10); i++ {
		maxWord := GetMaxWord(sameFreqWords).Word

		sortedWords = append(sortedWords, maxWord)

		sameFreqWords = deleteWord(sameFreqWords, maxWord)
	}

	return sortedWords, lenSFWs
}

func Top10(text string) []string {
	// Разбиваем текст на слова, разделяемые пробелами

	words := strings.Fields(text)

	// Словарь для подсчета частоты каждого слова

	wordCounts := make(map[string]int)

	for _, word := range words {
		wordCounts[strings.ToLower(word)]++
	}

	// Создаем слайс для сортировки

	frequencyWords := make([]wordCount, 0, len(wordCounts))

	for word, count := range wordCounts {
		frequencyWords = append(frequencyWords, wordCount{Word: word, Count: count, Length: len(word)})
	}

	// Сортировка слайса по убыванию частоты и лексикографически

	sort.Slice(frequencyWords, func(i, j int) bool {
		if frequencyWords[i].Count != frequencyWords[j].Count {
			return frequencyWords[i].Count > frequencyWords[j].Count
		}

		return frequencyWords[i].Length < frequencyWords[j].Length
	})

	// Возвращаем первые 10 слов

	topWords := make([]string, 0, len(wordCounts))

	for i := 0; i < 10 && i < len(frequencyWords); i++ {
		// topWords[sortedWords[i].Word] = sortedWords[i].Count

		lexicSort, indPlus := sortSameFrequencyWord(frequencyWords, frequencyWords[i].Count, i)

		for j := 0; j < len(lexicSort); j++ {
			topWords = append(topWords, lexicSort[j])
		}

		i = i + indPlus - 1

		// topWords = append(topWords, sortedWords[i].Word)
	}

	return topWords
}

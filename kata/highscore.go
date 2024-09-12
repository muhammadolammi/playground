package kata

import "strings"

var store = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
}

func High(s string) string {
	// your code here
	words := strings.Fields(s)
	wordCounts := make([]int, len(words))
	for i, word := range words {
		wordCounts[i] = calculateWordCount(word)
	}
	maxIndex := findMaxIndex(wordCounts)
	return words[maxIndex]
}

func calculateWordCount(s string) int {
	count := 0
	for _, ru := range s {
		count += store[string(ru)]
	}
	return count
}

func findMaxIndex(counts []int) int {
	max := -1
	maxIndex := 0
	for i, count := range counts {
		if count > max {
			max = count
			maxIndex = i
		}
	}
	return maxIndex
}

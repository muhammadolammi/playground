package main

import (
	"strings"
)

func uncommonFromSentences(s1 string, s2 string) []string {
	records := map[string]int{}
	uncommons := []string{}
	words1 := strings.Split(s1, " ")
	words2 := strings.Split(s2, " ")

	for _, s := range words1 {
		records[s]++
	}
	for _, s := range words2 {
		records[s]++
	}

	for s, count := range records {
		if count == 1 {
			uncommons = append(uncommons, s)
		}
	}

	return uncommons
}

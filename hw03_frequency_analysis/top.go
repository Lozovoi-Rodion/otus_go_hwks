package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type WordOccurrence struct {
	word  string
	times int
}

func Top10(str string) []string {
	if len(str) == 0 {
		return []string{}
	}

	words := strings.Fields(str)

	occurrences := make(map[string]int)
	for _, w := range words {
		occurrences[w]++
	}

	wordOccurrences := make([]WordOccurrence, 0, len(occurrences))
	for word, times := range occurrences {
		wordOccurrences = append(wordOccurrences, WordOccurrence{word, times})
	}

	sort.Slice(wordOccurrences, func(i, j int) bool {
		if wordOccurrences[i].times == wordOccurrences[j].times {
			return wordOccurrences[i].word < wordOccurrences[j].word
		}
		return wordOccurrences[i].times > wordOccurrences[j].times
	})

	topLength := 10
	if len(wordOccurrences) > topLength {
		wordOccurrences = wordOccurrences[:topLength]
	}

	result := make([]string, 0, len(wordOccurrences))
	for _, w := range wordOccurrences {
		result = append(result, w.word)
	}

	return result
}

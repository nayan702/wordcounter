package utils

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

// WordCount holds the word and its frequency
type WordCount struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}

type Content struct {
	Title       string
	Heading     string
	Description string
}

func ReadLines(filename string) ([]string, error) {
	var lines []string
	file, err := os.Open(filename)
	if err != nil {
		return lines, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) > 0 {
			lines = append(lines, line)
		}
	}
	return lines, scanner.Err()
}

func CreateWordBankSet(wordBank []string) map[string]struct{} {
	wordBankSet := make(map[string]struct{})
	for _, word := range wordBank {
		wordBankSet[strings.ToLower(word)] = struct{}{}
	}
	return wordBankSet
}

func CountValidWords(words []string, wordBankSet map[string]struct{}) map[string]int {
	counts := make(map[string]int)
	for _, word := range words {
		wordLower := strings.ToLower(word)
		if len(wordLower) >= 3 && isAlpha(wordLower) {
			if _, exists := wordBankSet[wordLower]; exists {
				counts[wordLower]++
			}
		}
	}
	return counts
}

func isAlpha(s string) bool {
	for _, r := range s {
		if !('a' <= r && r <= 'z') && !('A' <= r && r <= 'Z') {
			return false
		}
	}
	return true
}

func GetTopWords(wordCounts map[string]int, topN int) []WordCount {
	wordCountList := make([]WordCount, 0, len(wordCounts))
	for word, count := range wordCounts {
		wordCountList = append(wordCountList, WordCount{Word: word, Count: count})
	}
	sort.Slice(wordCountList, func(i, j int) bool {
		return wordCountList[i].Count > wordCountList[j].Count
	})
	if len(wordCountList) > topN {
		return wordCountList[:topN]
	}
	return wordCountList
}

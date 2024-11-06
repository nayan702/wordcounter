package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"wordcounter/constants"
	"wordcounter/handler"
	"wordcounter/parser"
	"wordcounter/utils"

	"golang.org/x/time/rate"
)

func main() {
	urls_file := "urls.txt"
	if len(os.Args) != 1 {
		urls_file = os.Args[1]
	}
	urls, err := utils.ReadLines(urls_file)
	if err != nil {
		log.Fatalf("Error reading URLs: %v", err)
	}
	log.Printf("Read %d URLs from %s\n", len(urls), urls_file)

	wordBank, err := utils.ReadLines("bank.txt")
	if err != nil {
		log.Fatalf("Error reading word bank: %v", err)
	}
	log.Printf("Read %d words from bank.txt\n", len(wordBank))

	wordBankSet := utils.CreateWordBankSet(wordBank)

	limiter := rate.NewLimiter(constants.MaxConcurrentRequestsPerSec, constants.RateLimiterBurstSize)
	results := handler.FetchContents(urls, limiter)

	wordCounts := make(map[string]int)
	var mu sync.Mutex

	for result := range results {
		if result.Error != nil {
			log.Printf("Error fetching URL %s: %v\n", result.URL, result.Error)
			continue
		}

		words := parser.ExtractWords(result.Content)
		localCounts := utils.CountValidWords(words, wordBankSet)

		mu.Lock()
		for word, count := range localCounts {
			wordCounts[word] += count
		}
		mu.Unlock()
	}

	topWords := utils.GetTopWords(wordCounts, 10)

	output, err := json.MarshalIndent(topWords, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}
	fmt.Println(string(output))
}

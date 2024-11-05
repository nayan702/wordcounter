package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
	"wordcounter/utils"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/time/rate"
)

type FetchResult struct {
	Content utils.Content
	Error   error
	URL     string
}

// FetchContents fetches content from a list of URLs with dynamic rate limiting.
func FetchContents(urls []string, limiter *rate.Limiter) <-chan FetchResult {
	results := make(chan FetchResult, len(urls))
	var wg sync.WaitGroup
	var mu sync.Mutex          // Mutex for synchronizing access to the limiter
	limitConcurrentRequests := make(chan struct{}, int(limiter.Limit())) // Channel to limit concurrent requests

	originalLimit := limiter.Limit()  // Store the original limit
	successCount := 0                 // Count of successful requests

	for _, url := range urls {
		wg.Add(1)
		limitConcurrentRequests <- struct{}{} // Block if the limit is reached
		
		go func(url string) {
			defer wg.Done()

			// Acquire a slot to limit concurrent requests
			defer func() { <-limitConcurrentRequests }() // Release the slot

			
			// Retry logic with backoff
			for attempt := 0; attempt < 3; attempt++ { // Example with 3 attempts
				// Wait for the rate limiter
				if err := limiter.Wait(context.Background()); err != nil {
					log.Printf("Error waiting for limiter: %v", err)
					return
				}

				content, err := FetchURL(url)
				results <- FetchResult{Content: content, Error: err, URL: url}

				if err == nil {
					successCount++ // Increment on successful fetch
					adjustRate(limiter, originalLimit, successCount)
					break // Break the retry loop on success
				}

				log.Printf("Error fetching URL %s: %v (attempt %d)", url, err, attempt+1)

				// Adjust rate if we receive a 999 response
				if err != nil && strings.Contains(err.Error(), "adjust rate limit") {
					mu.Lock() // Lock while adjusting
					newLimit := limiter.Limit() / 2 // Halve the limit
					// Ensure the new limit is at least 1 request per second
					if newLimit >= 1 {
						log.Printf("Adjusting rate limit from %v to %v, because of %v", limiter.Limit(), limiter.Limit()/2, err)
						limiter.SetLimit(rate.Limit(newLimit)) // Set the new limit
					}
					mu.Unlock() // Unlock after adjustment
					time.Sleep(60 * time.Second) // Backoff for 60 seconds
				}  else if strings.Contains(err.Error(), "404") {
					// For a 404 error, we can log it and break the retry loop without blocking others
					log.Printf("URL %s returned a 404 error. Skipping this URL.", url)
					break // Exit retry loop for this URL
				} else {
					// Exponential backoff for other errors
					time.Sleep(time.Duration(attempt+1) * time.Second)
				}
			}
		}(url)
	}
	wg.Wait()
	close(results)
	return results
}

// FetchURL retrieves the content from the specified URL.
func FetchURL(url string) (utils.Content, error) {
	resp, err := http.Get(url)
	if err != nil {
		return utils.Content{}, fmt.Errorf("error fetching URL: %v", err)
	}
	defer resp.Body.Close()

	// Handle the 999/429/503 response code specifically
	if resp.StatusCode == 999 || resp.StatusCode == 429 || resp.StatusCode == 503 {
		return utils.Content{}, fmt.Errorf("received %d response code, adjust rate limit", resp.StatusCode)
	}
	if resp.StatusCode != 200 {
		return utils.Content{}, fmt.Errorf("received non-200 response code %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return utils.Content{}, fmt.Errorf("error loading HTTP response body: %v", err)
	}

	title := doc.Find("meta[property='og:title']").AttrOr("content", "Title not found")
	heading := doc.Find("div.caas-subheadline h2").Text()
	var description strings.Builder
	doc.Find("div.caas-body p").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		description.WriteString(text + "\n")
	})

	return utils.Content{Title: title, Heading: heading, Description: description.String()}, nil
}

// adjustRate increments the rate limit back to the original value after successful fetches.
func adjustRate(limiter *rate.Limiter, originalLimit rate.Limit, successCount int) {
	const successThreshold = 50 // Number of successes to trigger a limit increase

	if successCount%successThreshold == 0 && limiter.Limit() < originalLimit {
		newLimit := limiter.Limit() + (originalLimit / 10) // Increase limit by 10% of the original limit
		if newLimit > originalLimit {
			newLimit = originalLimit // Don't exceed the original limit
		}
		limiter.SetLimit(newLimit)
		log.Printf("Adjusted rate limit back to %v", newLimit)
	}
}
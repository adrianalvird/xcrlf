package scanner

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"crlfi/internal/utils"
)

func Scan(target string, verbose bool) ([]string, error) {
	payloads := GetPayloads()
	userAgents := GetUserAgents()
	var results []string

	for _, payload := range payloads {
		url := fmt.Sprintf("%s%s", target, payload)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			if verbose {
				fmt.Printf("Error creating request for %s: %v\n", url, err)
			}
			continue
		}

		// Rotate User-Agent
		req.Header.Set("User-Agent", userAgents[rand.Intn(len(userAgents))])

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			if verbose {
				fmt.Printf("Error sending request to %s: %v\n", url, err)
			}
			continue
		}
		defer resp.Body.Close()

		// Check response status
		if resp.StatusCode == http.StatusOK {
			results = append(results, fmt.Sprintf("Potential CRLF injection at: %s", url))
		}

		// Add random delay
		AddRandomDelay()
	}

	return results, nil
}

func ScanConcurrent(targets []string, workers int, verbose bool) []string {
	jobs := make(chan string, len(targets))
	results := make(chan string, len(targets))
	var wg sync.WaitGroup

	// Worker pool
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for target := range jobs {
				res, err := Scan(target, verbose)
				if err != nil {
					continue
				}
				for _, r := range res {
					results <- r
				}
			}
		}()
	}

	// Add targets to job queue
	for _, target := range targets {
		jobs <- target
	}
	close(jobs)

	// Wait for workers to finish
	wg.Wait()
	close(results)

	// Collect results
	var allResults []string
	for result := range results {
		allResults = append(allResults, result)
	}
	return allResults
}

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	"crlfi/internal/scanner"
	"crlfi/internal/utils"
)

func main() {
	resumeFile := flag.String("r", "", "Resume scan using resume configuration file")
	outputFile := flag.String("o", "", "Output file for results")
	workers := flag.Int("w", 10, "Number of concurrent workers")
	verbose := flag.Bool("v", false, "Enable verbose output")
	flag.Parse()

	// Read targets from stdin
	scannerInput := bufio.NewScanner(os.Stdin)
	var targets []string
	for scannerInput.Scan() {
		targets = append(targets, scannerInput.Text())
	}
	if len(targets) == 0 {
		fmt.Println("No targets provided. Provide targets via stdin.")
		return
	}

	// Resume scan if resume file is provided
	if *resumeFile != "" {
		targets = scanner.ResumeScan(*resumeFile, targets)
	}

	// Open output file if specified
	var output *os.File
	if *outputFile != "" {
		var err error
		output, err = os.OpenFile(*outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Failed to open output file: %v\n", err)
			return
		}
		defer output.Close()
	}

	// Perform scanning with concurrency
	results := scanner.ScanConcurrent(targets, *workers, *verbose)

	// Write results to output file if specified
	for _, result := range results {
		fmt.Println(result)
		if output != nil {
			output.WriteString(result + "\n")
		}
	}
}

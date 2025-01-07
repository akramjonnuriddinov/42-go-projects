package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	// Read text from a file
	fmt.Print("Enter the path of the text file: ")
	var filePath string
	fmt.Scanln(&filePath)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	// Read the file content
	scanner := bufio.NewScanner(file)
	var text string
	for scanner.Scan() {
		text += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	// Perform text analysis
	wordCount, commaCount, periodCount, maxSentenceLength, minSentenceLength, maxWordLength, minWordLength := analyzeText(text)

	// Output the results
	fmt.Println("Text Analysis Results:")
	fmt.Printf("Word Count: %d\n", wordCount)
	fmt.Printf("Comma Count: %d\n", commaCount)
	fmt.Printf("Period Count: %d\n", periodCount)
	fmt.Printf("Max Sentence Length: %d\n", maxSentenceLength)
	fmt.Printf("Min Sentence Length: %d\n", minSentenceLength)
	fmt.Printf("Max Word Length: %d\n", maxWordLength)
	fmt.Printf("Min Word Length: %d\n", minWordLength)
}

func analyzeText(text string) (int, int, int, int, int, int, int) {
	// Count words
	words := strings.Fields(text)
	wordCount := len(words)

	// Count commas and periods
	commaCount := strings.Count(text, ",")
	periodCount := strings.Count(text, ".")

	// Split text into sentences based on periods
	sentences := splitSentences(text)
	maxSentenceLength := 0
	minSentenceLength := -1
	for _, sentence := range sentences {
		sentenceLength := len(strings.Fields(sentence))
		if sentenceLength > maxSentenceLength {
			maxSentenceLength = sentenceLength
		}
		if minSentenceLength == -1 || sentenceLength < minSentenceLength {
			minSentenceLength = sentenceLength
		}
	}

	// Find the longest and shortest words
	maxWordLength := 0
	minWordLength := -1
	for _, word := range words {
		wordLength := len(word)
		if wordLength > maxWordLength {
			maxWordLength = wordLength
		}
		if minWordLength == -1 || wordLength < minWordLength {
			minWordLength = wordLength
		}
	}

	return wordCount, commaCount, periodCount, maxSentenceLength, minSentenceLength, maxWordLength, minWordLength
}

// Split text into sentences based on periods
func splitSentences(text string) []string {
	re := regexp.MustCompile(`[^.!?]+[.!?]*`)
	return re.FindAllString(text, -1)
}

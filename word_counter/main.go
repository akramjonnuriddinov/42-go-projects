package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your text: ")
	if scanner.Scan() {
		words := scanner.Text()
		wordCount := wordCounter(words)
		fmt.Println(wordCount)
	} else if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}

func wordCounter(words string) int {
	trimmedWords := strings.TrimSpace(words)
	if trimmedWords == "" {
		fmt.Println("Try again with words")
		return 0
	}
	return len(strings.Fields(trimmedWords))
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
    fmt.Println("Please enter some text:")
    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        input := scanner.Text()
        processedInput := strings.ReplaceAll(input, " ", "")

        charCount := len(processedInput)
        fmt.Printf("Number of characters (excluding spaces): %d\n", charCount)

        frequencyList := make(map[rune]int)
        for _, char := range processedInput {
            frequencyList[char]++
        }
		
		var chars []rune
		for char := range frequencyList {
			chars = append(chars, char)
		}
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars [j]
		})


        fmt.Println("Frequency list of all characters used:")
        for _, char := range chars {
            fmt.Printf("%c appears %d times\n", char, frequencyList[char])
        }
	
	}

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "Error reading from input:", err)
    }
}
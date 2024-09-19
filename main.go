package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <string>")
		return
	}

	firstString := os.Args[1]
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fileContent := string(data)
	fileLines := strings.Split(fileContent, "\n")

	// Prepare an array for each line of the final output
	result := make([]string, 8) // Assuming each character is represented by 8 lines

	for _, char := range firstString {
		if char < 32 || char > 126 {
			fmt.Printf("Error: Character '%c' is not in range\n", char)
			continue
		}

		asci := char - 32
		startIndex := int(asci)*8 + int(asci) + 1

		// Ensure we don't go out of bounds
		if startIndex+7 < len(fileLines) {
			for j := 0; j < 8; j++ {
				result[j] += fileLines[startIndex+j] + "  " // Add space between characters
			}
		} else {
			fmt.Println("Error: Not enough lines for character representation")
		}
	}

	// Print the final result
	for _, line := range result {
		fmt.Println(line)
	}
}

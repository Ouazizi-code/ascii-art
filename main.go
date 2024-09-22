package main

import (
	"ascii-art/functions"
	"fmt"
	"os"
)

func main() {
	// check the args befor extracting our args
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <string>")
		return
	}

	// extract data from the standard file as an array of caracters
	data := functions.ReadFile("standard.txt")

	// now extraxt our text from the args
	str := os.Args[1]

	functions.Traitment(str, data)
	//fmt.Println(result)
	/*
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
	*/
}

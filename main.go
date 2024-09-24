package main

import (
	"fmt"
	"os"

	"ascii-art/functions"
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
 
	// send this args to trairment and print inside this function
	functions.Traitment(str, data)
}

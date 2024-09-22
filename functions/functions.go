package functions

import (
	"fmt"
	"os"
	"strings"
)

// this func just to read the standard file
func ReadFile(filename string) []string {

	data, err := os.ReadFile(filename)
	// handle err
	if err != nil {
		fmt.Println("ereur : ", err)
	}

	result := strings.Split(string(data), "\n")
	return result

}

// this function for traitment
func Traitment(str string, data []string) {
	//result := ""
	arr := make([]string, 8)
	splited_string := strings.Split(str, `\n`)

	/////////////////////////////////
	status := 0
	for _, test := range splited_string {
		if test == "" {
			status++
		}
	}
	if status == (len(str)/2)+1 {
		for i := 0; i <(len(str)/2) ; i++ {
			fmt.Println("good")
		}
		return
	}
	//////////////////////////////////////

	// loop throught the splited string as an array of indexes
	for _, part := range splited_string {

		// now loop throught each part of splited string
		for i := 0; i < len(part); i++ {
			// now we have the first caracter in our args <char>
			char := part[i]

			// lets initialse a start index in data depend on the char
			asci := char - 32
			startIndex := int(asci)*8 + int(asci) + 1
			// lets check if our caracter is printable or not
			if char < 32 || char > 126 {
				fmt.Println("ereur : this caracter is not in range => ", "'", string(char), "'")
				os.Exit(1)
			} else {

				for j := 0; j < 8; j++ {
					arr[j] += data[startIndex+j]
				}

				//arr[8][len(part)*8] += byte("\n")

			}
		}

		// lets check if the array is empty
		count := 0
		for _, check := range arr {
			if check == "" {
				count++
			}
		}

		// if count == 8 that mean arr is empty
		// print just a newline
		if count == len(arr) {
			fmt.Println()
		} else {
			// else the array is full print the text inside
			for _, line := range arr {
				fmt.Println(line)
			}
		}
		// reset the array for the next part
		arr = make([]string, 8)

	}

}

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
func Traitment(str string, data []string) string {
	result := ""

	return result
}

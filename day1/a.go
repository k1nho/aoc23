package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// Read file argument
	args := os.Args
	// Get filename input
	if len(args) != 2 {
		panic("no file provided")
	}
	filename := args[1]
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		panic("could not open file")
	}
	defer file.Close()

	// get the file content
	fileContent, err := io.ReadAll(file)
	if err != nil {
		panic("could not read filecontent")
	}
	// transform into string
	data := string(fileContent)
	solve(data)

}

func solve(data string) {
	lines := strings.Split(data, "\n")
	// ignore last entry (empty line)
	lines = lines[:len(lines)-1]

	res := int64(0)
	for _, line := range lines {
		var digits []string
		for _, ch := range line {
			if unicode.IsDigit(ch) {
				digits = append(digits, fmt.Sprintf("%c", ch))
			}
		}
		val, err := strconv.ParseInt(digits[0]+digits[len(digits)-1], 10, 64)
		if err != nil {
			panic("no val")
		}
		res += val
	}

	fmt.Print("res: ", res)
}

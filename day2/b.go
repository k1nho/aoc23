package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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
	lines = lines[:len(lines)-1]
	res := int64(0)

	for id, line := range lines {
		res += gameRes(line, id+1)
	}
	fmt.Println("answer: ", res)

}

func gameRes(line string, id int) int64 {
	games := strings.Split(line, ";")
	r, g, b := int64(1), int64(1), int64(1)
	for _, game := range games {
		balls := strings.Split(game, ",")
		for _, ball := range balls {
			num, t := fmt.Sprintf("%s", ball[:len(ball)-1]), fmt.Sprintf("%c", ball[len(ball)-1])
			val, err := strconv.ParseInt(num, 10, 64)
			if err != nil {
				panic("")
			}
			switch t {
			case "r":
				r = maxInt(r, val)
			case "g":
				g = maxInt(g, val)
			case "b":
				b = maxInt(b, val)
			}
		}
	}
	return r * g * b
}

func maxInt(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

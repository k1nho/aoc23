package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func maxInt(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func stoi(a string) int64 {
	val, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("could not convert %s to int", a))
	}
	return val
}

func sizeChecker(a []string, num int) {
	if len(a) < num {
		panic(fmt.Sprintf("size of arr: %d, size required: %d", len(a), num))
	}
}

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
	time := stoi(lines[0])
	record := stoi(lines[1])

	fmt.Println("time ", time, " record ", record)

	fmt.Println(calc(time, record))
}

func calc(time, record int64) int64 {
	// total time - pressed time * left time

	ways := int64(0)

	// binary search is better here but this will do
	for i := int64(0); i < time; i++ {
		if (time-i)*i > record {
			ways += 1
		}
	}
	return ways
}

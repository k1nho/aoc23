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

	res := int64(0)
	for _, line := range lines {
		res += calc(line)
	}
	fmt.Println(res)
}

func calc(line string) int64 {
	res := int64(0)
	spl := strings.Split(line, "|")
	winningNums := strings.Split(strings.TrimSpace(spl[0]), " ")
	nums := strings.Split(strings.TrimSpace(spl[1]), " ")

	st := map[string]bool{}
	for _, num := range winningNums {
		st[num] = true
	}

	// check which ones match and double
	for _, num := range nums {
		if num == "" {
			continue
		}
		if _, ok := st[num]; ok {
			if res == 0 {
				res = 1
			} else {
				res *= 2
			}
		}
	}

	return res
}

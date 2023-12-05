package main

import (
	"fmt"
	"io"
	"math"
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
	lines := strings.Split(data, "\n\n")
	lines = lines[:]

	seedsString := lines[0]
	seeds := []int64{}

	for _, seed := range strings.Split(seedsString, " ") {
		seeds = append(seeds, stoi(seed))
	}

	fmt.Println(seeds)

	lines = lines[1:]

	conversionMaps := [][][]int64{}
	for _, line := range lines {
		conversionMaps = append(conversionMaps, calc(line))
	}

	res := int64(math.MaxInt64)
	for _, seed := range seeds {
		temp := seed
		for _, mp := range conversionMaps {
			for _, v := range mp {
				dest, source, to := v[0], v[1], v[2]
				fmt.Println("is temp ", temp, " in range ", source, source+to)
				if temp >= source && temp <= source+to-1 {
					temp = dest + (temp - source)
					break
				}
			}
		}
		res = Min(res, temp)
	}
	fmt.Println(res)

}

func calc(line string) [][]int64 {
	spl := strings.Split(line, "\n")
	arr := [][]int64{}

	factors := spl[0]
	facts := strings.Split(factors, " ")
	fmt.Println(facts[0], facts[1])
	for i := 1; i < len(spl); i++ {
		ranges := strings.Split(spl[i], " ")
		if len(ranges) == 0 {
			continue
		}
		if ranges[0] == "" || ranges[1] == "" || ranges[2] == "" {
			continue
		}
		dest, source, to := stoi(ranges[0]), stoi(ranges[1]), stoi(ranges[2])
		arr = append(arr, []int64{dest, source, to})

	}
	return arr
}

func Min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	directions = [][]int{{-1, 0}, {-1, 1}, {-1, -1}, {1, 0}, {1, 1}, {1, -1}, {0, -1}, {0, 1}}
	symbols    = map[string]bool{"@": true, "#": true, "$": true, "%": true, "&": true, "*": true, "-": true, "+": true, "/": true, "=": true}
)

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func isCoordinateInBounds(x, y, xMax, yMax int) bool {
	return x >= 0 && x < xMax && y >= 0 && y < yMax
}

func maxInt(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func convertInt(a string) int64 {
	val, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("could not convert %s into int\n", a))
	}
	return val
}

func isInt(a string) bool {
	if _, err := strconv.Atoi(a); err == nil {
		return true
	}
	return false
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
	graph := [][]string{}
	for _, line := range lines {
		arr := strings.Split(line, "")
		graph = append(graph, arr)
	}
	fmt.Println(calc(graph))
}

func calc(graph [][]string) int64 {
	res := int64(0)
	r, c := len(graph), len(graph[0])

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if symbols[graph[i][j]] {
				for _, p := range directions {
					nr, nc := p[0]+i, p[1]+j
					if isCoordinateInBounds(nr, nc, r, c) && isInt(graph[nr][nc]) {
						res += getNum(graph, nr, nc, c)
					}
				}
			}
		}
	}
	return res
}

func getNum(graph [][]string, nr, nc, c int) int64 {
	num := ""
	left, right := nc-1, nc+1

	// walk left
	for (left >= 0) && isInt(graph[nr][left]) {
		num = graph[nr][left] + num
		graph[nr][left] = "."
		left--
	}

	num = num + graph[nr][nc]
	graph[nr][nc] = "."

	//walk right
	for (right < c) && isInt(graph[nr][right]) {
		num = num + graph[nr][right]
		graph[nr][right] = "."
		right++
	}

	return convertInt(num)
}

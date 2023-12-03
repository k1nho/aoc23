package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	top    = [][]int{{-1, 0}, {-1, 1}, {-1, -1}}
	bottom = [][]int{{1, 0}, {1, 1}, {1, -1}}
	lr     = [][]int{{0, -1}, {0, 1}}
)

func isCoordinateInBounds(x, y, xMax, yMax int) bool {
	return x >= 0 && x < xMax && y >= 0 && y < yMax
}

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
			if graph[i][j] == "*" {
				adj := [][]int{}
				for _, p := range top {
					nr, nc := p[0]+i, p[1]+j
					if isCoordinateInBounds(nr, nc, r, c) && isInt(graph[nr][nc]) {
						adj = append(adj, []int{nr, nc})
						if p[1] == 0 {
							break
						}
					}
				}
				for _, p := range bottom {
					nr, nc := p[0]+i, p[1]+j
					if isCoordinateInBounds(nr, nc, r, c) && isInt(graph[nr][nc]) {
						adj = append(adj, []int{nr, nc})
						if p[1] == 0 {
							break
						}
					}
				}
				for _, p := range lr {
					nr, nc := p[0]+i, p[1]+j
					if isCoordinateInBounds(nr, nc, r, c) && isInt(graph[nr][nc]) {
						adj = append(adj, []int{nr, nc})
					}
				}

				// can there be a gear with more than 2 adjacent numbers?
				if len(adj) == 2 {
					total := int64(1)
					for _, ad := range adj {
						total *= getNum(graph, ad[0], ad[1], c)
					}
					res += total
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
		left--
	}

	num = num + graph[nr][nc]

	//walk right
	for (right < c) && isInt(graph[nr][right]) {
		num = num + graph[nr][right]
		right++
	}

	return convertInt(num)
}

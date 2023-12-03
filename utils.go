package main

import (
	"fmt"
	"strconv"
)

// GENERAL

var (
	symbols = map[string]bool{"@": true, "#": true, "$": true, "%": true, "&": true, "*": true, "-": true, "+": true, "/": true, "=": true}
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

// GRAPHS
var (
	directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, 1}, {-1, -1}, {1, 1}, {1, -1}}
	nseo       = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	top        = [][]int{{-1, 0}, {-1, 1}, {-1, -1}}
	bottom     = [][]int{{1, 0}, {1, 1}, {1, -1}}
	lr         = [][]int{{0, -1}, {0, 1}}
)

func isCoordinateInBounds(x, y, xMax, yMax int) bool {
	return x >= 0 && x < xMax && y >= 0 && y < yMax
}

package main

import (
	"fmt"
	"io"
	"os"
	"sort"
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
	cards := [][]string{}

	for _, line := range lines {
		spl := strings.Split(line, " ")
		cards = append(cards, []string{spl[0], spl[1]})
	}
	mp := map[string][][]string{}
	for _, card := range cards {
		mp[getRank(card[0])] = append(mp[getRank(card[0])], card)
	}

	// sort the ranks starting from five
	total := int64(0)
	rankNum := int64(len(lines))
	rankTypes := []string{"five", "four", "full", "three", "two", "one", "high"}
	for _, rankType := range rankTypes {
		total += sortRank(mp[rankType], &rankNum)

	}
	fmt.Println("total ", total)

}

func sortRank(rank [][]string, rankNum *int64) int64 {
	letterToVal := map[string]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		"J": 1,
		"T": 10,
	}
	if len(rank) == 0 {
		return 0
	}

	sort.Slice(rank, func(i, j int) bool {
		a, b := rank[i][0], rank[j][0]
		powera, powerb := 0, 0
		for i := 0; i < len(a); i++ {
			achar := fmt.Sprintf("%c", a[i])
			bchar := fmt.Sprintf("%c", b[i])
			if val, ok := letterToVal[achar]; ok {
				powera = val
			} else {
				powera = int(stoi(achar))
			}
			if val, ok := letterToVal[bchar]; ok {
				powerb = val
			} else {
				powerb = int(stoi(bchar))
			}

			if powera == powerb {
				continue
			} else {
				return powera > powerb
			}
		}
		return powera > powerb
	})

	ans := int64(0)
	for _, card := range rank {
		ans += stoi(card[1]) * (*rankNum)
		*rankNum -= 1
	}
	return ans
}

func getRank(card string) string {
	// Five of a kind, where all five cards have the same label: AAAAA
	// Four of a kind, where four cards have the same label and one card has a different label: AA8AA
	// Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
	// Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
	// Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
	// One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
	// High card, where all cards' labels are distinct: 23456

	freqs := map[string]int64{}
	seenJs := 0
	for _, ch := range card {
		str := fmt.Sprintf("%c", ch)
		if str == "J" {
			seenJs += 1
		} else {
			freqs[str] += 1
		}
	}

	for key := range freqs {
		freqs[key] += int64(seenJs)
	}
	if seenJs == 5 {
		freqs["J"] = 5
	}

	switch len(freqs) {
	case 1:
		return "five"
	case 2:
		// can be four, full house
		seenFour := false
		for _, val := range freqs {
			if val == 4 {
				seenFour = true
			}
		}
		if seenFour {
			return "four"
		} else {
			return "full"
		}

	case 3:
		// can be three, two pair
		seenThree := false
		for _, val := range freqs {
			if val == 3 {
				seenThree = true
			}
		}
		if seenThree {
			return "three"
		} else {
			return "two"
		}

	case 4:
		// one pair
		return "one"
	default:
		return "high"
	}

}

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var goTmpl = `package main

import (
	"io"
	"os"
)

func check(err error) {
    if err != nil {
        panic(err.Error())
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

}
`

var makefileTmpl = `# GOLANG
# input test
gd:
	go run a.go i1.txt
# sample test
gs:
	go run a.go i2.txt
# input test part 2
gd2:
	go run b.go i1.txt
# sample test part 2
gs2:
	go run b.go i2.txt
`

func main() {
	args := os.Args
	if len(args) != 2 {
		panic("no day number provided")
	}
	daynum := args[1]
	userDir, err := os.UserHomeDir()
	if err != nil {
		panic(err.Error())
	}

	// create folder
	dayDir := fmt.Sprintf("%s/%s%s", userDir, "Developer/aoc2023/day", filepath.Base(daynum))
	if err = os.Mkdir(dayDir, os.ModePerm); err != nil {
		panic(fmt.Sprintf("could not create day folder: %s", err.Error()))
	}

	// create files
	filenames := []string{"a", "b"}
	for _, filename := range filenames {
		file, err := os.Create(filepath.Join(dayDir, fmt.Sprintf("%s.go", filename)))
		if err != nil {
			panic("could not create file")
		}
		if _, err = file.WriteString(goTmpl); err != nil {
			panic("could not write file")
		}
		file.Close()
	}

	// input files
	filenames = []string{"i1", "i2"}
	for _, filename := range filenames {
		file, err := os.Create(filepath.Join(dayDir, fmt.Sprintf("%s.txt", filename)))
		if err != nil {
			panic("could not create file")
		}
		file.Close()
	}

	// makefile
	file, err := os.Create(filepath.Join(dayDir, "Makefile"))
	if err != nil {
		panic("could not create makefile")
	}

	if _, err = file.WriteString(makefileTmpl); err != nil {
		panic("could not write makefile")
	}
	file.Close()

	fmt.Printf("files created for day %s of AOC\n", daynum)
}

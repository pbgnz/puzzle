package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/pbgnz/puzzle/puzzle"
)

func main() {

	inputs := handleInput(os.Args[1:])

	fmt.Println("Running dfs")
	p := puzzle.NewPuzzle(inputs)
	dfs := puzzle.DepthFirstSearch(p)
	generateOutputFiles(dfs, "output/puzzleDFS.txt")

	fmt.Println("Running bfsH1")
	puzzle2 := puzzle.NewPuzzle(inputs)
	bfsH1 := puzzle.BestFirstSearch(puzzle2, 1)
	generateOutputFiles(bfsH1, "output/puzzleBFS-h1.txt")

	fmt.Println("Running bfsH2")
	puzzle3 := puzzle.NewPuzzle(inputs)
	bfsH2 := puzzle.BestFirstSearch(puzzle3, 2)
	generateOutputFiles(bfsH2, "output/puzzleBFS-h2.txt")

	fmt.Println("Running asH1")
	puzzle4 := puzzle.NewPuzzle(inputs)
	asH1 := puzzle.As(puzzle4, 1)
	generateOutputFiles(asH1, "output/puzzleAs-h1.txt")

	fmt.Println("Running asH2")
	puzzle5 := puzzle.NewPuzzle(inputs)
	asH2 := puzzle.As(puzzle5, 2)
	generateOutputFiles(asH2, "output/puzzleAs-h2.txt")

}

// handleInput takes the user input (array of strings),
// validates it, and returns an array of integers
func handleInput(i []string) []int {
	if len(i) != 12 {
		fmt.Println("Error: you must enter 12 integers [0,11]")
		os.Exit(1)
	}
	ints := make([]int, 12)
	inputExists := map[int]bool{}
	for i, v := range i {
		input, _ := strconv.Atoi(v)

		if input < 0 || input > 11 {
			fmt.Println("Error: input(s) not in domain [0,11]")
			os.Exit(1)
		}

		if inputExists[input] {
			fmt.Println("Error: there are duplicate inputs")
			os.Exit(1)
		}

		inputExists[input] = true
		ints[i] = input
	}
	return ints
}

// generateOutputFiles prints the goal path in a given file
func generateOutputFiles(p []*puzzle.Node, l string) {
	s := ""
	for i := len(p) - 1; i >= 0; i-- {
		s += p[i].Move + " [" + strings.Trim(strings.Replace(fmt.Sprint(p[i].Puzzle), " ", " , ", -1), "[]") + "]\n"
		ioutil.WriteFile(l, []byte(s), 0666)
	}
}

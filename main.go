package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/pbgnz/11d-puzzle/puzzle"
)

func main() {

	args := os.Args[1:]
	if len(args) != 12 {
		fmt.Println("Wrong inputs")
	}

	ints := make([]int, 12)
	for i, v := range args {
		ints[i], _ = strconv.Atoi(v)
	}
	p := puzzle.NewPuzzle(ints)
	r := puzzle.BreadthFirstSearch(p)
	generateOutputFiles(r)

}

func generateOutputFiles(p []*puzzle.Node) {
	s := ""
	for i := len(p) - 1; i >= 0; i-- {
		s += p[i].Move + " [" + strings.Trim(strings.Replace(fmt.Sprint(p[i].Puzzle), " ", " , ", -1), "[]") + "]\n"
		ioutil.WriteFile("output/output.txt", []byte(s), 0666)
	}
}

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

	args := os.Args[1:]
	if len(args) != 12 {
		fmt.Println("Wrong inputs")
	}

	ints := make([]int, 12)
	for i, v := range args {
		ints[i], _ = strconv.Atoi(v)
	}
	p := puzzle.NewPuzzle(ints)
	dfs := puzzle.DepthFirstSearch(p)
	generateOutputFiles(dfs, "output/puzzleDFS.txt")

	bfs := puzzle.BreadthFirstSearch(p)
	generateOutputFiles(bfs, "output/puzzleBFS.txt")

}

func generateOutputFiles(p []*puzzle.Node, l string) {
	s := ""
	for i := len(p) - 1; i >= 0; i-- {
		s += p[i].Move + " [" + strings.Trim(strings.Replace(fmt.Sprint(p[i].Puzzle), " ", " , ", -1), "[]") + "]\n"
		ioutil.WriteFile(l, []byte(s), 0666)
	}
}

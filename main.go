package main

import (
	"fmt"
	"os"

	"github.com/pbgnz/11d-puzzle/puzzle"
)

func main() {

	args := os.Args[1:]
	if len(args) != 12 {
		fmt.Println("Wrong inputs")
	}

	p := puzzle.NewPuzzle(args)
	for _, e := range p.Puzzle {
		fmt.Println(e)
	}
	p.MoveUp(args, 4)

	for _, e := range p.Children[0].Puzzle {
		fmt.Println(e)
	}
}

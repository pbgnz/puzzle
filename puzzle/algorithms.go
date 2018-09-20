package puzzle

import (
	"fmt"
)

// DepthFirstSearch prints
func DepthFirstSearch(r *Node) []*Node {
	o := make([]*Node, 0) // open-list (stack)
	c := make([]*Node, 0) // closed-list
	p := make([]*Node, 0) // solution path
	foundPath := false

	o = append(o, r)

	for len(o) > 0 && !foundPath {
		current := o[0]
		o = o[1:]

		fmt.Println(current.Puzzle)
		if current.isGoalState() {
			foundPath = true
			p = current.PathTrace()
		}
		c = append([]*Node{current}, c...)
		current.GenerateMoves()

		for i := 0; i < len(current.Children); i++ {
			if !Contains(o, current.Children[i]) && !Contains(c, current.Children[i]) {
				o = append(o, current.Children[i])
			}
		}
	}
	return p
}

// BreadthFirstSearch prints
func BreadthFirstSearch(r *Node) []*Node {
	o := make([]*Node, 0) // open-list (queue)
	c := make([]*Node, 0) // closed-list
	p := make([]*Node, 0) // solution path
	foundPath := false

	o = append(o, r)

	for len(o) > 0 && !foundPath {
		current := o[0]
		o = o[1:]
		c = append(c, current)

		current.GenerateMoves()

		for i := 0; i < len(current.Children); i++ {

			x := current.Children[i]

			if x.isGoalState() {
				fmt.Println("Goal found")
				foundPath = true
				p = x.PathTrace()
			}

			if !Contains(o, x) && !Contains(c, x) {
				o = append(o, x)
			}
		}

	}
	return p
}

func Contains(s []*Node, n *Node) bool {
	contains := false
	for i := 0; i < len(s); i++ {
		if AreTheSame(s[i].Puzzle, n.Puzzle) {
			contains = true
		}
	}
	return contains
}

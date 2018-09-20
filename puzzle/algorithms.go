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
		current, o := o[0], o[1:]

		if current.isGoalState() {
			foundPath = true
			c = append(c, current)
			return PathTrace(current)
		}
		current.GenerateMoves()

		for i := 0; i < len(current.Children); i++ {

			x := current.Children[i]
			if !Contains(o, x) && !Contains(c, x) {
				o = append([]*Node{x}, o...)
			}
		}
		c = append(c, current)

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
		current, o := o[0], o[1:]
		c = append(c, current)

		current.GenerateMoves()

		for i := 0; i < len(current.Children); i++ {

			x := current.Children[i]

			if x.isGoalState() {
				fmt.Println("Goal found")
				foundPath = true
				p = PathTrace(x)
			}

			if !Contains(o, x) && !Contains(c, x) {
				o = append(o, x)
			}
		}

	}
	return p
}

// Contains as
func Contains(s []*Node, n *Node) bool {
	for _, e := range s {
		if e == n {
			return true
		}
	}
	return false
}

func PathTrace(e *Node) []*Node {
	n := make([]*Node, 0)
	current := e
	n = append(n, current)
	fmt.Println(n)
	for current.Parent != nil {
		current = current.Parent
		n = append(n, current)
		fmt.Println(n)
	}
	return n
}

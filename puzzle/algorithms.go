package puzzle

// DepthFirstSearch algorithm
func DepthFirstSearch(r *Node) []*Node {
	openList := make([]*Node, 0) //(stack)
	closedList := make([]*Node, 0)
	solutionPath := make([]*Node, 0)
	foundPath := false

	// add the root node to the openList
	openList = append(openList, r)

	for len(openList) > 0 && !foundPath {

		// remove left-most element from openList
		// call it x
		x := openList[0]
		openList = openList[1:]

		// check if x is goal
		if x.isGoalState() {
			foundPath = true
			solutionPath = x.PathTrace()
		}

		// genrate children of x
		x.GenerateMoves()

		// put x on closedList
		closedList = append([]*Node{x}, closedList...)

		for i := 0; i < len(x.Children); i++ {
			if !Contains(openList, x.Children[i]) && !Contains(closedList, x.Children[i]) {
				// put remaining children of x on left end of open list
				openList = append([]*Node{x.Children[i]}, openList...)
			}
		}
	}
	return solutionPath
}

// BreadthFirstSearch algorithm
func BreadthFirstSearch(r *Node) []*Node {
	openList := make([]*Node, 0) // (queue)
	closedList := make([]*Node, 0)
	solutionPath := make([]*Node, 0)
	foundPath := false

	openList = append(openList, r)

	for len(openList) > 0 && !foundPath {
		x := openList[0]
		openList = openList[1:]
		closedList = append(closedList, x)

		if x.isGoalState() {
			foundPath = true
			solutionPath = x.PathTrace()
		}

		x.GenerateMoves()

		for i := 0; i < len(x.Children); i++ {

			child := x.Children[i]

			if !Contains(openList, child) && !Contains(closedList, child) {
				// put remaining children of x on right end of open list
				openList = append(openList, child)
			}
		}

	}
	return solutionPath
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

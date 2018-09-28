package puzzle

import (
	"container/heap"
	"fmt"
)

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

// BestFirstSearch algorithm
func BestFirstSearch(root *Node, heuristic int) []*Node {
	openList := make(PriorityQueue, 1)
	closedList := make([]*Node, 0)
	solutionPath := make([]*Node, 0)
	foundPath := false

	// add the root node to the openList
	if heuristic == 1 {
		root.Heuristic = Heuristic1(root.Puzzle)
	} else {
		root.Heuristic = Heuristic2(root.Puzzle)
	}
	item := &Item{
		Value:    root,
		Priority: root.Heuristic,
	}
	openList[0] = item
	heap.Init(&openList)

	fmt.Println("here")
	for openList.Len() > 0 && !foundPath {

		x := openList.Pop().(*Item).Value
		closedList = append(closedList, x)

		if x.isGoalState() {
			foundPath = true
			solutionPath = x.PathTrace()
		}

		x.GenerateMoves()

		for i := 0; i < len(x.Children); i++ {
			child := x.Children[i]
			if !Contains(closedList, child) {
				// add the root node to the openList
				if heuristic == 1 {
					child.Heuristic = Heuristic1(child.Puzzle)
				} else {
					child.Heuristic = Heuristic2(child.Puzzle)
				}
				item := &Item{
					Value:    child,
					Priority: child.Heuristic,
				}
				// put remaining children of x on left end of open list
				openList.Push(item)
				heap.Push(&openList, item)
			}
		}

	}
	return solutionPath
}

// As algorithm
func As(root *Node, heuristic int) []*Node {
	openList := make(PriorityQueue, 1)
	closedList := make([]*Node, 0)
	solutionPath := make([]*Node, 0)
	foundPath := false

	// add the root node to the openList
	if heuristic == 1 {
		root.Heuristic = Heuristic1(root.Puzzle) + root.G
	} else {
		root.Heuristic = Heuristic2(root.Puzzle) + root.G
	}
	item := &Item{
		Value:    root,
		Priority: root.Heuristic,
	}
	openList[0] = item
	heap.Init(&openList)

	fmt.Println("here")
	for openList.Len() > 0 && !foundPath {

		x := openList.Pop().(*Item).Value
		closedList = append(closedList, x)

		if x.isGoalState() {
			foundPath = true
			solutionPath = x.PathTrace()
		}

		x.GenerateMoves()

		for i := 0; i < len(x.Children); i++ {
			child := x.Children[i]
			if !Contains(closedList, child) {
				// add the root node to the openList
				if heuristic == 1 {
					child.Heuristic = Heuristic1(child.Puzzle) + child.G
				} else {
					child.Heuristic = Heuristic2(child.Puzzle) + child.G
				}
				item := &Item{
					Value:    child,
					Priority: child.Heuristic,
				}
				// put remaining children of x on left end of open list
				openList.Push(item)
				heap.Push(&openList, item)
			}
		}

	}
	return solutionPath
}

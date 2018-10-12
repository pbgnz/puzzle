package puzzle

import (
	"container/heap"
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
				openList = append(openList, x.Children[i])
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
		Node:      root,
		Heuristic: root.Heuristic,
	}
	openList[0] = item
	heap.Init(&openList)

	for openList.Len() > 0 && !foundPath {

		x := heap.Pop(&openList).(*Item).Node
		closedList = append(closedList, x)

		if x.isGoalState() {
			foundPath = true
			solutionPath = x.PathTrace()
		}

		x.GenerateMoves()

		for i := 0; i < len(x.Children); i++ {
			if x.Children[i].isGoalState() {
				foundPath = true
				solutionPath = x.Children[i].PathTrace()
			} else {
				if heuristic == 1 {
					x.Children[i].Heuristic = Heuristic1(x.Children[i].Puzzle)
				} else {
					x.Children[i].Heuristic = Heuristic2(x.Children[i].Puzzle)
				}
			}
		}

		for i := 0; i < len(x.Children); i++ {
			childFoundInOpen, oldChildOpen := openList.Contains(x.Children[i])
			childFoundInClose, oldChildClose := ContainsAndRemove(&closedList, x.Children[i])

			if !childFoundInOpen && !childFoundInClose {
				item := &Item{
					Node:      x.Children[i],
					Heuristic: x.Children[i].Heuristic,
				}
				heap.Push(&openList, item)
			} else if childFoundInOpen {
				if x.Children[i].Heuristic < oldChildOpen.Node.Heuristic {
					oldChildOpen.Node.Heuristic = x.Children[i].Heuristic
					openList.update(oldChildOpen, oldChildOpen.Node, oldChildOpen.Node.Heuristic)
				}
			} else if childFoundInClose {
				item := &Item{
					Node:      oldChildClose,
					Heuristic: oldChildClose.Heuristic,
				}
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
		Node:      root,
		Heuristic: root.Heuristic,
	}
	openList[0] = item
	heap.Init(&openList)

	for openList.Len() > 0 && !foundPath {

		x := heap.Pop(&openList).(*Item).Node
		closedList = append(closedList, x)

		if x.isGoalState() {
			foundPath = true
			solutionPath = x.PathTrace()
		}

		x.GenerateMoves()

		for i := 0; i < len(x.Children); i++ {
			if x.Children[i].isGoalState() {
				foundPath = true
				solutionPath = x.Children[i].PathTrace()
			} else {
				if heuristic == 1 {
					x.Children[i].Heuristic = Heuristic1(x.Children[i].Puzzle) + x.Children[i].G
				} else {
					x.Children[i].Heuristic = Heuristic2(x.Children[i].Puzzle) + x.Children[i].G
				}
			}
		}

		for i := 0; i < len(x.Children); i++ {
			childFoundInOpen, oldChildOpen := openList.Contains(x.Children[i])
			childFoundInClose, oldChildClose := ContainsAndRemove(&closedList, x.Children[i])

			if !childFoundInOpen && !childFoundInClose {
				item := &Item{
					Node:      x.Children[i],
					Heuristic: x.Children[i].Heuristic,
				}
				heap.Push(&openList, item)
			} else if childFoundInOpen {
				if x.Children[i].G < oldChildOpen.Node.G {
					oldChildOpen.Node.Parent = x.Children[i].Parent
					openList.update(oldChildOpen, oldChildOpen.Node, oldChildOpen.Node.Heuristic)
				}
			} else if childFoundInClose {
				if x.Children[i].G < oldChildClose.G {
					oldChildClose.Parent = x.Children[i].Parent
					item := &Item{
						Node:      oldChildClose,
						Heuristic: oldChildClose.Heuristic,
					}
					heap.Push(&openList, item)
				} else {
					item := &Item{
						Node:      oldChildClose,
						Heuristic: oldChildClose.Heuristic,
					}
					heap.Push(&openList, item)
				}
			}
		}
	}
	return solutionPath
}

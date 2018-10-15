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

	// until openList is empty or a goal node has been reached
	for openList.Len() > 0 && !foundPath {

		// remove from openList a node at which f is minimum
		// call it x
		x := heap.Pop(&openList).(*Item).Node
		// put x on closedList
		closedList = append(closedList, x)

		if x.isGoalState() {
			foundPath = true
			solutionPath = x.PathTrace()
		}

		// genrate children of x
		x.GenerateMoves()

		// if any of the children of x is a goal node, exit.
		for i := 0; i < len(x.Children); i++ {
			if x.Children[i].isGoalState() {
				foundPath = true
				solutionPath = x.Children[i].PathTrace()
			} else {
				// calculate f(child)
				if heuristic == 1 {
					x.Children[i].Heuristic = Heuristic1(x.Children[i].Puzzle)
				} else {
					x.Children[i].Heuristic = Heuristic2(x.Children[i].Puzzle)
				}
			}
		}

		// for each child of x
		for i := 0; i < len(x.Children); i++ {
			// check if child was on openList or closedList
			// results used below
			childFoundInOpen, oldChildOpen := openList.Contains(x.Children[i])
			childFoundInClose, oldChildClose := ContainsAndRemove(&closedList, x.Children[i])

			// if child was neither on openList nor closedList, add it to
			// openList.  Attach the cost f(child) to it.
			if !childFoundInOpen && !childFoundInClose {
				item := &Item{
					Node:      x.Children[i],
					Heuristic: x.Children[i].Heuristic,
				}
				heap.Push(&openList, item)
			} else if childFoundInOpen {
				// 	if child was on the openList; compare
				//  the new value f(child) with the previously assigned
				//  value f(oldChildOpen).  If the old value is lower discard
				//  the newly generated node.  If the new value is
				//  lower, substitute the new node for the old node.
				if x.Children[i].Heuristic < oldChildOpen.Node.Heuristic {
					// modify the priority and value of an Item in the queue.
					// the node is already in the priority queue
					oldChildOpen.Node.Heuristic = x.Children[i].Heuristic
					openList.update(oldChildOpen, oldChildOpen.Node, oldChildOpen.Node.Heuristic)
				}
			} else if childFoundInClose {
				// 	If the old node was on closedList move it back to
				//  openList.
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

	// until openList is empty or a goal node has been reached
	for openList.Len() > 0 && !foundPath {

		// remove from openList a node at which f is minimum
		// call it x
		x := heap.Pop(&openList).(*Item).Node
		// put x on closedList
		closedList = append(closedList, x)

		// check if x is goal
		if x.isGoalState() {
			foundPath = true
			solutionPath = x.PathTrace()
		}

		// genrate children of x
		x.GenerateMoves()

		for i := 0; i < len(x.Children); i++ {
			if x.Children[i].isGoalState() {
				foundPath = true
				solutionPath = x.Children[i].PathTrace()
			} else {
				// calculate f(child)
				if heuristic == 1 {
					x.Children[i].Heuristic = Heuristic1(x.Children[i].Puzzle) + x.Children[i].G
				} else {
					x.Children[i].Heuristic = Heuristic2(x.Children[i].Puzzle) + x.Children[i].G
				}
			}
		}

		for i := 0; i < len(x.Children); i++ {
			// check if child was on openList or closedList
			// results used below
			childFoundInOpen, oldChildOpen := openList.Contains(x.Children[i])
			childFoundInClose, oldChildClose := ContainsAndRemove(&closedList, x.Children[i])

			// if child was neither on openList nor closedList, add
			// it to openList. Attach the cost f(child) to it.
			if !childFoundInOpen && !childFoundInClose {
				item := &Item{
					Node:      x.Children[i],
					Heuristic: x.Children[i].Heuristic,
				}
				heap.Push(&openList, item)
			} else if childFoundInOpen {
				// if child was already on openList
				// direct its pointers along the path yielding
				// the lowest g(child) and keep the lowest f(child).
				if x.Children[i].G < oldChildOpen.Node.G {
					// modify the priority and value of an Item in the queue
					oldChildOpen.Node.Parent = x.Children[i].Parent
					openList.update(oldChildOpen, oldChildOpen.Node, oldChildOpen.Node.Heuristic)
				}
			} else if childFoundInClose {
				// 	If the old node was on closedList move it back to
				//  openList.
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

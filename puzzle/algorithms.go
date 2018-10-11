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
		Value:    root,
		Priority: root.Heuristic,
	}
	openList[0] = item
	heap.Init(&openList)

	// until openList is empty or a goal node has been reached
	for openList.Len() > 0 && !foundPath {

		// remove from openList a node at which f is minimum
		// call it x
		x := openList.Pop().(*Item).Value

		// put x on closedList
		closedList = append(closedList, x)

		// check if x is goal
		if x.isGoalState() {
			foundPath = true
			solutionPath = x.PathTrace()
		}

		// genrate children of x
		x.GenerateMoves()

		// if any of the children of x is a goal node, exit.
		for i := 0; i < len(x.Children); i++ {
			// check if child is goal
			if x.Children[i].isGoalState() {
				foundPath = true
				solutionPath = x.Children[i].PathTrace()
			}
		}

		// for each child of x
		for i := 0; i < len(x.Children); i++ {
			child := x.Children[i]

			// check if child was on openList or closedList
			// results used below
			inOpenList, oldChildOpen := openList.Contains(child)
			inClosedList, oldChildClosed := ContainsAndRemove(&closedList, child)

			if inClosedList && inOpenList {
				fmt.Println("Something wrong")
			}
			// calculate f(child)
			if heuristic == 1 {
				child.Heuristic = Heuristic1(child.Puzzle)
			} else {
				child.Heuristic = Heuristic2(child.Puzzle)
			}

			// if child was neither on openList nor closedList, add it to
			// openList.  Attach the cost f(child) to it.
			if !inClosedList && !inOpenList {
				item := &Item{
					Value:    child,
					Priority: child.Heuristic,
				}
				openList.Push(item)
				heap.Push(&openList, item)
			} else if inOpenList {
				// 	if child was on the openList; compare
				//  the new value f(child) with the previously assigned
				//  value f(oldChildOpen).  If the old value is lower discard
				//  the newly generated node.  If the new value is
				//  lower, substitute the new node for the old node.
				if child.Heuristic < oldChildOpen.Heuristic {
					oldChildOpen.Heuristic = child.Heuristic
					// modify the priority and value of an Item in the queue.
					// the node is already in the priority queue
					openList.update(&Item{
						Value:    oldChildOpen,
						Priority: oldChildOpen.Heuristic,
					}, oldChildOpen, oldChildOpen.Heuristic)
				}
			} else if inClosedList {
				// 	If the old node was on closedList move it back to
				//  openList.
				item := &Item{
					Value:    oldChildClosed,
					Priority: oldChildClosed.Heuristic,
				}
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

	// until openList is empty or a goal node has been reached
	for openList.Len() > 0 && !foundPath {

		// remove from openList a node at which f is minimum
		// call it x
		x := openList.Pop().(*Item).Value

		// put x on closedList
		closedList = append(closedList, x)

		// check if x is goal
		if x.isGoalState() {
			foundPath = true
			solutionPath = x.PathTrace()
		}

		// genrate children of x
		x.GenerateMoves()

		// foreach child of x
		for i := 0; i < len(x.Children); i++ {
			child := x.Children[i]

			// check if child was on openList or closedList
			// results used below
			inOpenList, oldChildOpen := openList.Contains(child)
			inClosedList, oldChildClosed := ContainsAndRemove(&closedList, child)
			if inClosedList && inOpenList {
				fmt.Println("Something wrong")
			}
			// calculate f(child)
			if heuristic == 1 {
				child.Heuristic = Heuristic1(child.Puzzle) + child.G
			} else {
				child.Heuristic = Heuristic2(child.Puzzle) + child.G
			}

			// if child was neither on openList nor closedList, add
			// it to openList. Attach the cost f(child) to it.
			if !inClosedList && !inOpenList {
				item := &Item{
					Value:    child,
					Priority: child.Heuristic,
				}
				openList.Push(item)
				heap.Push(&openList, item)
			} else if inOpenList {
				// if child was already on openList
				// direct its pointers along the path yielding
				// the lowest g(child) and keep the lowest f(child).
				if child.G < oldChildOpen.G {
					oldChildOpen.G = child.G
					oldChildOpen.Parent = child.Parent
					if child.Heuristic < oldChildOpen.Heuristic {
						oldChildOpen.Heuristic = child.Heuristic
					}
					// modify the priority and value of an Item in the queue.
					openList.update(&Item{
						Value:    oldChildOpen,
						Priority: oldChildOpen.Heuristic,
					}, oldChildOpen, oldChildOpen.Heuristic)
				} else {
					if child.Heuristic < oldChildOpen.Heuristic {
						oldChildOpen.Heuristic = child.Heuristic
					}
					// modify the priority and value of an Item in the queue.
					openList.update(&Item{
						Value:    oldChildOpen,
						Priority: oldChildOpen.Heuristic,
					}, oldChildOpen, oldChildOpen.Heuristic)
				}
			} else if inClosedList {
				// 	If the old node was on closedList move it back to
				//  openList.
				if child.G < oldChildClosed.G {
					oldChildClosed.G = child.G
					oldChildClosed.Parent = child.Parent
					if child.Heuristic < oldChildClosed.Heuristic {
						oldChildClosed.Heuristic = child.Heuristic
					}
				} else {
					if child.Heuristic < oldChildClosed.Heuristic {
						oldChildClosed.Heuristic = child.Heuristic
					}
				}
				item := &Item{
					Value:    oldChildClosed,
					Priority: oldChildClosed.Heuristic,
				}
				openList.Push(item)
				heap.Push(&openList, item)
			}
		}

	}
	return solutionPath
}

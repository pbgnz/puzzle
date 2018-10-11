package puzzle

import "reflect"

// Node represents a in the search tree
type Node struct {
	Puzzle    []int
	Value     int // the index of the "0" in the puzzle
	Children  []*Node
	Parent    *Node
	Move      string // used for printing files
	Heuristic int
	G         int // the depth of the node in the search tree
}

// NewPuzzle generates the root node
func NewPuzzle(p []int) *Node {
	puzzle := make([]int, NumberColumns*NumberRows)
	for i := 0; i < len(p); i++ {
		puzzle[i] = p[i]
	}
	return &Node{
		Puzzle:    puzzle,
		Value:     -1,
		Children:  make([]*Node, 0),
		Parent:    nil,
		Move:      "0",
		Heuristic: -1,
		G:         0,
	}
}

// GenerateMoves generates the possible moves
func (n *Node) GenerateMoves() {
	for i := 0; i < len(n.Puzzle); i++ {
		if n.Puzzle[i] == 0 {
			n.Value = i
		}
	}

	n.MoveUp(n.Puzzle, n.Value)
	n.MoveUpRight(n.Puzzle, n.Value)
	n.MoveRight(n.Puzzle, n.Value)
	n.MoveDownRight(n.Puzzle, n.Value)
	n.MoveDown(n.Puzzle, n.Value)
	n.MoveDownLeft(n.Puzzle, n.Value)
	n.MoveLeft(n.Puzzle, n.Value)
	n.MoveUpLeft(n.Puzzle, n.Value)
}

// GoalTest veries if a puzzle is the goal state
func (n *Node) isGoalState() bool {
	for i := 0; i < len(n.Puzzle)-1; i++ {
		if n.Puzzle[i] != goalState[i] {
			return false
		}
	}
	return true
}

// PathTrace returns a slice of nodes leading to goal
func (n *Node) PathTrace() []*Node {
	new := make([]*Node, 0)
	current := n
	new = append(new, current)
	for current.Parent != nil {
		current = current.Parent
		new = append(new, current)
	}
	return new
}

// ClonePuzzle copies the puzzle of a given Node
func (n *Node) ClonePuzzle() []int {
	p := make([]int, NumberColumns*NumberRows)
	for i := 0; i < len(n.Puzzle); i++ {
		p[i] = n.Puzzle[i]
	}
	return p
}

// AreTheSame compares two slices
func AreTheSame(a []int, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Contains checks if a slice of Nodes contains a given Node
func Contains(s []*Node, n *Node) bool {
	for i := 0; i < len(s); i++ {
		if AreTheSame(s[i].Puzzle, n.Puzzle) {
			if reflect.DeepEqual(s[i], n) {
				return true
			}
		}
	}
	return false
}

// ContainsAndRemove checks if an slice of Nodes contains a given Node
// and removes the node if it is theres
func ContainsAndRemove(s *[]*Node, n *Node) (bool, *Node) {
	for i := 0; i < len(*s); i++ {
		if AreTheSame((*s)[i].Puzzle, n.Puzzle) {
			if reflect.DeepEqual((*s)[i], n) {
				node := (*s)[i]
				(*s) = append((*s)[:i], (*s)[i+1:]...)
				return true, node
			}
		}
	}
	return false, nil
}

// MoveUp moves the empty tile up by 1 tile
func (n *Node) MoveUp(p []int, i int) {
	if i-NumberColumns >= 0 {
		c := n.ClonePuzzle()
		temp := c[i-NumberColumns]
		c[i-NumberColumns] = c[i]
		c[i] = temp

		child := NewPuzzle(c)
		child.Move = boardPositions[i-NumberColumns]
		child.Parent = n
		child.G = n.G + 1
		n.Children = append(n.Children, child)
	}
}

// MoveUpRight moves the empty tile diagonally up-right by 1 tile
func (n *Node) MoveUpRight(p []int, i int) {
	if i%NumberColumns != 3 && i > 3 {
		c := n.ClonePuzzle()
		temp := c[i-NumberRows]
		c[i-NumberRows] = c[i]
		c[i] = temp

		child := NewPuzzle(c)
		child.Move = boardPositions[i-NumberRows]
		child.Parent = n
		child.G = n.G + 1
		n.Children = append(n.Children, child)
	}
}

// MoveRight moves the empty tile to the right by 1 tile
func (n *Node) MoveRight(p []int, i int) {
	if i%NumberColumns != 3 {
		c := n.ClonePuzzle()
		temp := c[i+1]
		c[i+1] = c[i]
		c[i] = temp

		child := NewPuzzle(c)
		child.Move = boardPositions[i+1]
		child.Parent = n
		child.G = n.G + 1
		n.Children = append(n.Children, child)
	}
}

// MoveDownRight moves the empty tile diagonally down-right by 1 tile
func (n *Node) MoveDownRight(p []int, i int) {
	if i%NumberColumns != 3 && i < 8 {
		c := n.ClonePuzzle()
		temp := c[i+5]
		c[i+5] = c[i]
		c[i] = temp

		child := NewPuzzle(c)
		child.Move = boardPositions[i+5]
		child.Parent = n
		child.G = n.G + 1
		n.Children = append(n.Children, child)
	}
}

// MoveDown moves the empty tile down by 1 tile
func (n *Node) MoveDown(p []int, i int) {
	if i < 8 {
		c := n.ClonePuzzle()
		temp := c[i+4]
		c[i+4] = c[i]
		c[i] = temp

		child := NewPuzzle(c)
		child.Move = boardPositions[i+4]
		child.Parent = n
		child.G = n.G + 1
		n.Children = append(n.Children, child)
	}
}

// MoveDownLeft moves the empty tile diagonally down-left by 1 tile
func (n *Node) MoveDownLeft(p []int, i int) {
	if i%NumberColumns > 0 && i < 8 {
		c := n.ClonePuzzle()
		temp := c[i+3]
		c[i+3] = c[i]
		c[i] = temp

		child := NewPuzzle(c)
		child.Move = boardPositions[i+3]
		child.Parent = n
		child.G = n.G + 1
		n.Children = append(n.Children, child)
	}
}

// MoveLeft moves the empty tile left by 1 tile
func (n *Node) MoveLeft(p []int, i int) {
	if i%NumberColumns > 0 {
		c := n.ClonePuzzle()
		temp := c[i-1]
		c[i-1] = c[i]
		c[i] = temp

		child := NewPuzzle(c)
		child.Move = boardPositions[i-1]
		child.Parent = n
		child.G = n.G + 1
		n.Children = append(n.Children, child)
	}
}

// MoveUpLeft moves the empty tile diagonally up-left by 1 tile
func (n *Node) MoveUpLeft(p []int, i int) {
	if i%NumberColumns > 0 && i-NumberColumns >= 0 {
		c := n.ClonePuzzle()
		temp := c[i-5]
		c[i-5] = c[i]
		c[i] = temp

		child := NewPuzzle(c)
		child.Move = boardPositions[i-5]
		child.Parent = n
		child.G = n.G + 1
		n.Children = append(n.Children, child)
	}
}

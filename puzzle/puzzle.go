package puzzle

const (
	// NumberRows of the board
	NumberRows = 3
	// NumberColumns of the board
	NumberColumns = 4
)

var boardPositions = map[int]string{
	0:  "a",
	1:  "b",
	2:  "c",
	3:  "d",
	4:  "e",
	5:  "f",
	6:  "g",
	7:  "h",
	8:  "i",
	9:  "j",
	10: "k",
	11: "l",
}

// Node represents a tree node
type Node struct {
	Value     int
	Children  []*Node
	Parent    *Node
	Puzzle    []int
	Move      string
	Heuristic int
}

// NewPuzzle generates the root node
func NewPuzzle(p []int) *Node {
	puzzle := make([]int, NumberColumns*NumberRows)
	for i := 0; i < len(p); i++ {
		puzzle[i] = p[i]
	}
	return &Node{
		Value:     -1,
		Children:  make([]*Node, 0),
		Parent:    nil,
		Puzzle:    puzzle,
		Move:      "0",
		Heuristic: -1,
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

// GoalTest ehnd
func (n *Node) isGoalState() bool {
	for i := 0; i < len(n.Puzzle)-1; i++ {
		if n.Puzzle[i] != i+1 {
			return false
		}
	}
	return n.Puzzle[len(n.Puzzle)-1] == 0
}

// PathTrace he
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

// ClonePuzzle d
func (n *Node) ClonePuzzle() []int {
	p := make([]int, NumberColumns*NumberRows)
	for i := 0; i < len(n.Puzzle); i++ {
		p[i] = n.Puzzle[i]
	}
	return p
}

// AreTheSame compare two slices
func AreTheSame(a []int, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// MoveUp moves the empty tile up by 1 tile
// @param {[]int} p - the puzzle (1 dimensional array)
// @param {int} i - the position of the empty tile
func (n *Node) MoveUp(p []int, i int) {
	if i-NumberColumns >= 0 {
		c := n.ClonePuzzle()
		temp := c[i-NumberColumns]
		c[i-NumberColumns] = c[i]
		c[i] = temp

		child := NewPuzzle(c)
		child.Move = boardPositions[i-NumberColumns]
		child.Parent = n
		n.Children = append(n.Children, child)
	}
}

// MoveUpRight moves the empty tile diagonally up-right by 1 tile
// @param {[]int} p - the puzzle (1 dimensional array)
// @param {int} i - the position of the empty tile
func (n *Node) MoveUpRight(p []int, i int) {
	if i%NumberColumns != 3 && i > 3 {
		c := n.ClonePuzzle()
		temp := c[i-NumberRows]
		c[i-NumberRows] = c[i]
		c[i] = temp

		child := NewPuzzle(c)
		child.Move = boardPositions[i-NumberRows]
		child.Parent = n
		n.Children = append(n.Children, child)
	}
}

// MoveRight moves the empty tile to the right by 1 tile
// @param {[]int} p - the puzzle (1 dimensional array)
// @param {int} i - the position of the empty tile
func (n *Node) MoveRight(p []int, i int) {
	if i%NumberColumns != 3 {
		c := n.ClonePuzzle()
		temp := c[i+1]
		c[i+1] = c[i]
		c[i] = temp

		child := NewPuzzle(c)
		child.Move = boardPositions[i+1]
		child.Parent = n
		n.Children = append(n.Children, child)
	}
}

// MoveDownRight moves the empty tile diagonally down-right by 1 tile
// @param {[]int} p - the puzzle (1 dimensional array)
// @param {int} i - the position of the empty tile
func (n *Node) MoveDownRight(p []int, i int) {
	if i%NumberColumns != 3 && i < 8 {
		c := n.ClonePuzzle()
		temp := c[i+5]
		c[i+5] = c[i]
		c[i] = temp

		child := NewPuzzle(c)
		child.Move = boardPositions[i+5]
		child.Parent = n
		n.Children = append(n.Children, child)
	}
}

// MoveDown moves the empty tile down by 1 tile
// @param {[]int} p - the puzzle (1 dimensional array)
// @param {int} i - the position of the empty tile
func (n *Node) MoveDown(p []int, i int) {
	if i < 8 {
		c := n.ClonePuzzle()
		temp := c[i+4]
		c[i+4] = c[i]
		c[i] = temp

		child := NewPuzzle(c)
		child.Move = boardPositions[i+4]
		child.Parent = n
		n.Children = append(n.Children, child)
	}
}

// MoveDownLeft moves the empty tile diagonally down-left by 1 tile
// @param {[]int} p - the puzzle (1 dimensional array)
// @param {int} i - the position of the empty tile
func (n *Node) MoveDownLeft(p []int, i int) {
	if i%NumberColumns > 0 && i < 8 {
		c := n.ClonePuzzle()
		temp := c[i+3]
		c[i+3] = c[i]
		c[i] = temp

		child := NewPuzzle(c)
		child.Move = boardPositions[i+3]
		child.Parent = n
		n.Children = append(n.Children, child)
	}
}

// MoveLeft moves the empty tile left by 1 tile
// @param {[]int} p - the puzzle (1 dimensional array)
// @param {int} i - the position of the empty tile
func (n *Node) MoveLeft(p []int, i int) {
	if i%NumberColumns > 0 {
		c := n.ClonePuzzle()
		temp := c[i-1]
		c[i-1] = c[i]
		c[i] = temp

		child := NewPuzzle(c)
		child.Move = boardPositions[i-1]
		child.Parent = n
		n.Children = append(n.Children, child)
	}
}

// MoveUpLeft moves the empty tile diagonally up-left by 1 tile
// @param {[]int} p - the puzzle (1 dimensional array)
// @param {int} i - the position of the empty tile
func (n *Node) MoveUpLeft(p []int, i int) {
	if i%NumberColumns > 0 && i-NumberColumns >= 0 {
		c := n.ClonePuzzle()
		temp := c[i-5]
		c[i-5] = c[i]
		c[i] = temp

		child := NewPuzzle(c)
		child.Move = boardPositions[i-5]
		child.Parent = n
		n.Children = append(n.Children, child)
	}
}

package puzzle

const (
	NumberColumns = 4
	NumberRows    = 3
)

// Node a
type Node struct {
	Value    int
	Children []*Node
	Parent   *Node
	Puzzle   []string
}

// NewPuzzle generates the root node
func NewPuzzle(p []string) *Node {
	puzzle := make([]string, NumberColumns*NumberRows)
	for i := 0; i < len(p); i++ {
		puzzle[i] = p[i]
	}
	return &Node{
		Value:    0,
		Children: make([]*Node, 0),
		Parent:   nil,
		Puzzle:   puzzle,
	}
}

// GoalTest ehnd
func (n *Node) GoalTest() bool {

	isGoal := true
	m := n.Puzzle[0]

	for i := 1; i < len(n.Puzzle); i++ {
		if m > n.Puzzle[i] {
			isGoal = false
		}
		m = n.Puzzle[i]
	}
	return isGoal
}

func (n *Node) MoveUp(p []string, i int) {
	if i-NumberColumns >= 0 {
		c := n.ClonePuzzle()
		temp := c[i-NumberColumns]
		c[i-NumberColumns] = c[i]
		c[i] = temp

		child := NewPuzzle(c)
		n.Children = append(n.Children, child)
		child.Parent = n
	}
}

func (n *Node) ClonePuzzle() []string {
	p := make([]string, NumberColumns*NumberRows)
	for i := 0; i < len(n.Puzzle); i++ {
		p[i] = n.Puzzle[i]
	}
	return p
}

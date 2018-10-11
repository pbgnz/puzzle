package puzzle

import (
	"reflect"
	"testing"
)

func TestNewPuzzle(t *testing.T) {
	i := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 0, 11}
	p := NewPuzzle(i)

	if len(p.Puzzle) != 12 {
		t.Errorf("Expected puzzle length of 12, but got %v", len(p.Puzzle))
	}

	for index, value := range p.Puzzle {
		if value != i[index] {
			t.Errorf("The puzzle did not match the inputs")
		}
	}
}

func TestGenerateMovesAndAreTheSame(t *testing.T) {
	i := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 0, 11}
	p := NewPuzzle(i)

	//  1 | 2 | 3 | 4
	//	5 | 6 | 7 | 8
	//	9 | 10| 0 | 11

	if len(p.Children) != 0 {
		t.Errorf("Expected puzzle not to have children, but got %v", len(p.Children))
	}

	// the puzzle has 5 possible moves
	// the 0 can go UP, UP-RIGHT, RIGHT
	// LEFT, UP-LEFT
	p.GenerateMoves()

	if len(p.Children) != 5 {
		t.Errorf("Expected puzzle to have 5 children, but got %v", len(p.Children))
	}

	// the children are added to the children slice
	// in order of priority, the highest priority elements
	// will be placed at the begining of the slice
	// UP > UP-RIGHT > RIGHT > DOWN-RIGHT >
	// DOWN > DOWN-LEFT > LEFT > UP-LEFT
	up := []int{1, 2, 3, 4, 5, 6, 0, 8, 9, 10, 7, 11}
	upRight := []int{1, 2, 3, 4, 5, 6, 7, 0, 9, 10, 8, 11}
	right := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 0}
	left := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 10, 11}
	upLeft := []int{1, 2, 3, 4, 5, 0, 7, 8, 9, 10, 6, 11}

	if !AreTheSame(p.Children[0].Puzzle, up) {
		t.Errorf("Expected the UP move to be generated first")
	}

	if !AreTheSame(p.Children[1].Puzzle, upRight) {
		t.Errorf("Expected the UP-RIGHT move to be generated second")
	}

	if !AreTheSame(p.Children[2].Puzzle, right) {
		t.Errorf("Expected the RIGHT move to be generated third")
	}

	if !AreTheSame(p.Children[3].Puzzle, left) {
		t.Errorf("Expected the LEFT move to be generated fourth")
	}

	if !AreTheSame(p.Children[4].Puzzle, upLeft) {
		t.Errorf("Expected the UP-LEFT move to be generated last")
	}
}

func TestIsGoalState(t *testing.T) {
	i := []int{5, 3, 7, 0, 11, 1, 6, 10, 4, 9, 2, 8}
	p := NewPuzzle(i)

	if p.isGoalState() != false {
		t.Errorf("Expected p not to be the goal state")
	}

	i2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 0, 11}
	p2 := NewPuzzle(i2)

	if p2.isGoalState() != false {
		t.Errorf("Expected p2 not to be the goal state")
	}

	i3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 0}
	p3 := NewPuzzle(i3)

	if p3.isGoalState() != true {
		t.Errorf("Expected p3 to be the goal state")
	}
}

func TestPathTrace(t *testing.T) {
	i := []int{1, 2, 3, 0, 5, 6, 7, 4, 9, 10, 11, 8}
	p := NewPuzzle(i)
	//  1 | 2 | 3 | 0
	//	5 | 6 | 7 | 4
	//	9 | 10| 11| 8

	p.GenerateMoves()

	p.Children[0].GenerateMoves()

	r := p.Children[0].Children[1].PathTrace()

	// results should return 2 moves + initial state
	if len(r) != 3 {
		t.Errorf("Expected the solution path to be size 3, but was size %v", len(r))
	}
	move1 := []int{1, 2, 3, 4, 5, 6, 7, 0, 9, 10, 11, 8}
	//  1 | 2 | 3 | 4
	//	5 | 6 | 7 | 0
	//	9 | 10| 11| 8

	move2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 0}
	//  1 | 2 | 3 | 4
	//	5 | 6 | 7 | 8
	//	9 | 10| 11| 0

	// [move2 (goal state), move1, intial state]
	if AreTheSame(r[0].Puzzle, move2) != true {
		t.Errorf("Expected the solution path to return the second move")
	}

	if AreTheSame(r[1].Puzzle, move1) != true {
		t.Errorf("Expected the solution path to return the first move")
	}

	if AreTheSame(r[2].Puzzle, i) != true {
		t.Errorf("Expected the solution path to return the intial state")
	}

}

func TestClonePuzzle(t *testing.T) {
	i := []int{1, 2, 3, 0, 5, 6, 7, 4, 9, 10, 11, 8}
	p := NewPuzzle(i)

	if AreTheSame(p.ClonePuzzle(), i) != true {
		t.Errorf("The puzzle was not cloned properly")
	}
}

func TestAreTheSame(t *testing.T) {

	s1 := []int{1, 2, 3, 0, 5, 6, 7, 4, 9, 10, 11, 8}
	s2 := []int{2, 1, 3, 0, 5, 2, 7, 4, 9, 10, 11, 8}
	s3 := []int{2, 1, 3, 0, 5, 2, 7, 4, 9, 10, 11, 8}

	if AreTheSame(s1, s1) != true {
		t.Errorf("The two slices are the same")
	}

	if AreTheSame(s2, s3) != true {
		t.Errorf("The two slices are the same")
	}

	if AreTheSame(s2, s2) != true {
		t.Errorf("The two slices are the same")
	}

	if AreTheSame(s1, s2) != false {
		t.Errorf("The two slices are different")
	}

	if AreTheSame(s1, s3) != false {
		t.Errorf("The two slices are different")
	}
}

func TestContains(t *testing.T) {
	s := make([]*Node, 0)
	i := []int{5, 3, 7, 0, 11, 1, 6, 10, 4, 9, 2, 8}
	p := NewPuzzle(i)
	p.GenerateMoves()

	b1 := Contains(s, p)
	if b1 != false {
		t.Errorf("Found a node in an empty slice")
	}

	s = append(s, p)
	b2 := Contains(s, p)
	if b2 != true {
		t.Errorf("Did not found the node on the slice")
	}

	s = append(s, p.Children[0])
	b3 := Contains(s, p.Children[0])
	if b3 != true {
		t.Errorf("Did not found the node on the slice")
	}

	b4 := Contains(p.Children, p.Children[0])
	if b4 != true {
		t.Errorf("Did not found the node on the slice")
	}

	b5 := Contains(p.Children, p.Children[1])
	if b5 != true {
		t.Errorf("Did not found the node on the slice")
	}

	b6 := Contains(p.Children, p.Children[2])
	if b6 != true {
		t.Errorf("Did not found the node on the slice")
	}

	if len(p.Children) != 3 {
		t.Errorf("Removed some nodes from the slice: %v remain", len(p.Children))
	}
}

func TestContainsAndRemove(t *testing.T) {
	s := make([]*Node, 0)
	i := []int{5, 3, 7, 0, 11, 1, 6, 10, 4, 9, 2, 8}
	p := NewPuzzle(i)
	p.GenerateMoves()

	b1, n1 := ContainsAndRemove(&s, p)
	if b1 != false {
		t.Errorf("Found a node in an empty slice")
	}
	if n1 != nil {
		t.Errorf("Found a node in an empty slice")
	}

	s = append(s, p.Children[0])
	if len(s) != 1 {
		t.Errorf("Node was not added to the slice")
	}
	b2, n2 := ContainsAndRemove(&s, p.Children[0])
	if b2 != true {
		t.Errorf("Did not found the node on the slice")
	}
	if !reflect.DeepEqual(p.Children[0], n2) {
		t.Errorf("Did not found the node on the slice")
	}
	if len(s) != 0 {
		t.Errorf("Node was not removed from the slice")
	}

	if len(p.Children) != 3 {
		t.Errorf("Suposed to have 3 children: %v children found", len(p.Children))
	}

	b4, _ := ContainsAndRemove(&p.Children, p.Children[0])
	if b4 != true {
		t.Errorf("Did not found the node on the slice")
	}

	if len(p.Children) != 2 {
		t.Errorf("Suposed to have 2 children: %v children found", len(p.Children))
	}

	b5, _ := ContainsAndRemove(&p.Children, p.Children[0])
	if b5 != true {
		t.Errorf("Did not found the node on the slice")
	}

	if len(p.Children) != 1 {
		t.Errorf("Suposed to have 1 children: %v children found", len(p.Children))
	}

	b6, _ := ContainsAndRemove(&p.Children, p.Children[0])
	if b6 != true {
		t.Errorf("Did not found the node on the slice")
	}

	if len(p.Children) != 0 {
		t.Errorf("Did not remove all nodes from the slice: %v remain", len(p.Children))
	}
}

func TestMoveUp(t *testing.T) {
	i := []int{1, 2, 3, 4, 5, 6, 0, 8, 9, 10, 7, 11}
	p := NewPuzzle(i)

	//  1 | 2 | 3 | 4
	//	5 | 6 | 0 | 8
	//	9 | 10| 7 | 11

	// find the position of the "0"
	// this is done in the GenerateMoves()
	for i := 0; i < len(p.Puzzle); i++ {
		if p.Puzzle[i] == 0 {
			p.Value = i
		}
	}

	p.MoveUp(p.Puzzle, p.Value)
	up := []int{1, 2, 0, 4, 5, 6, 3, 8, 9, 10, 7, 11}
	if AreTheSame(p.Children[0].Puzzle, up) != true {
		t.Errorf("MoveUp did not move the puzzle correctly")
	}

	// TEST INVALID MOVES

	// p.Children[0].Puzzle looks like this:
	//  1 | 2 | 0 | 4
	//	5 | 6 | 3 | 8
	//	9 | 10| 7 | 11
	// let's try to make an invalid move, and
	// move it up
	p.Children[0].Value = 2
	p.Children[0].MoveUp(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveUp did an invalid move at puzzle index 2")
	}

	//  0 | ? | ? | ?
	//	? | ? | ? | ?
	//	? | ? | ? | ?
	p.Children[0].Value = 0
	p.Children[0].MoveUp(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveUp did an invalid move at puzzle index 0")
	}

	//  ? | 0 | ? | ?
	//	? | ? | ? | ?
	//	? | ? | ? | ?
	p.Children[0].Value = 1
	p.Children[0].MoveUp(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveUp did an invalid move at puzzle index 1")
	}

	//  ? | ? | ? | 0
	//	? | ? | ? | ?
	//	? | ? | ? | ?
	p.Children[0].Value = 3
	p.Children[0].MoveUp(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveUp did an invalid move at puzzle index 3")
	}
}

func TestMoveUpRight(t *testing.T) {
	i := []int{1, 2, 3, 4, 5, 6, 0, 8, 9, 10, 7, 11}
	p := NewPuzzle(i)

	//  1 | 2 | 3 | 4
	//	5 | 6 | 0 | 8
	//	9 | 10| 7 | 11

	// find the position of the "0"
	// this is done in the GenerateMoves()
	for i := 0; i < len(p.Puzzle); i++ {
		if p.Puzzle[i] == 0 {
			p.Value = i
		}
	}

	p.MoveUpRight(p.Puzzle, p.Value)
	upRight := []int{1, 2, 3, 0, 5, 6, 4, 8, 9, 10, 7, 11}
	if AreTheSame(p.Children[0].Puzzle, upRight) != true {
		t.Errorf("MoveUpRight did not move the puzzle correctly")
	}

	// TEST INVALID MOVES

	// p.Children[0].Puzzle looks like this:
	//  1 | 2 | 3 | 0
	//	5 | 6 | 4 | 8
	//	9 | 10| 7 | 11
	// let's try to make an invalid move, and
	// move it up right
	p.Children[0].Value = 2
	p.Children[0].MoveUpRight(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveUpRight did an invalid move at puzzle index 2")
	}

	//  0 | ? | ? | ?
	//	? | ? | ? | ?
	//	? | ? | ? | ?
	p.Children[0].Value = 0
	p.Children[0].MoveUpRight(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveUpRight did an invalid move at puzzle index 0")
	}

	//  ? | 0 | ? | ?
	//	? | ? | ? | ?
	//	? | ? | ? | ?
	p.Children[0].Value = 1
	p.Children[0].MoveUpRight(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveUpRight did an invalid move at puzzle index 1")
	}

	//  ? | ? | ? | 0
	//	? | ? | ? | ?
	//	? | ? | ? | ?
	p.Children[0].Value = 3
	p.Children[0].MoveUpRight(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveUpRight did an invalid move at puzzle index 3")
	}

	//  ? | ? | ? | ?
	//	? | ? | ? | 0
	//	? | ? | ? | ?
	p.Children[0].Value = 7
	p.Children[0].MoveUpRight(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveUpRight did an invalid move at puzzle index 7")
	}

	//  ? | ? | ? | ?
	//	? | ? | ? | ?
	//	? | ? | ? | 0
	p.Children[0].Value = 11
	p.Children[0].MoveUpRight(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveUpRight did an invalid move at puzzle index 11")
	}

}

func TestMoveRight(t *testing.T) {
	i := []int{1, 2, 3, 4, 5, 6, 0, 8, 9, 10, 7, 11}
	p := NewPuzzle(i)

	//  1 | 2 | 3 | 4
	//	5 | 6 | 0 | 8
	//	9 | 10| 7 | 11

	// find the position of the "0"
	// this is done in the GenerateMoves()
	for i := 0; i < len(p.Puzzle); i++ {
		if p.Puzzle[i] == 0 {
			p.Value = i
		}
	}

	p.MoveRight(p.Puzzle, p.Value)
	right := []int{1, 2, 3, 4, 5, 6, 8, 0, 9, 10, 7, 11}
	if AreTheSame(p.Children[0].Puzzle, right) != true {
		t.Errorf("MoveRight did not move the puzzle correctly")
	}

	// TEST INVALID MOVES

	// p.Children[0].Puzzle looks like this:
	//  1 | 2 | 3 | 4
	//	5 | 6 | 8 | 0
	//	9 | 10| 7 | 11
	// let's try to make an invalid move, and
	// move it right
	p.Children[0].Value = 7
	p.Children[0].MoveRight(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveRight did an invalid move at puzzle index 2")
	}

	//  ? | ? | ? | 0
	//	? | ? | ? | ?
	//	? | ? | ? | ?
	p.Children[0].Value = 3
	p.Children[0].MoveRight(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveRight did an invalid move at puzzle index 3")
	}

	//  ? | ? | ? | ?
	//	? | ? | ? | ?
	//	? | ? | ? | 0
	p.Children[0].Value = 11
	p.Children[0].MoveRight(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveRight did an invalid move at puzzle index 11")
	}

}

func TestMoveDownRight(t *testing.T) {
	i := []int{1, 2, 3, 4, 5, 6, 0, 8, 9, 10, 7, 11}
	p := NewPuzzle(i)

	//  1 | 2 | 3 | 4
	//	5 | 6 | 0 | 8
	//	9 | 10| 7 | 11

	// find the position of the "0"
	// this is done in the GenerateMoves()
	for i := 0; i < len(p.Puzzle); i++ {
		if p.Puzzle[i] == 0 {
			p.Value = i
		}
	}

	p.MoveDownRight(p.Puzzle, p.Value)
	right := []int{1, 2, 3, 4, 5, 6, 11, 8, 9, 10, 7, 0}
	if AreTheSame(p.Children[0].Puzzle, right) != true {
		t.Errorf("MoveDownRight did not move the puzzle correctly")
	}

	// p.Children[0].Puzzle looks like this:
	//  1 | 2 | 3 | 4
	//	5 | 6 | 11| 8
	//	9 | 10| 7 | 0
	// let's try to make an invalid move, and
	// move it down right
	p.Children[0].Value = 11
	p.Children[0].MoveDownRight(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveDownRight did an invalid move at puzzle index 11")
	}

	//  ? | ? | ? | 0
	//	? | ? | ? | ?
	//	? | ? | ? | ?
	p.Children[0].Value = 3
	p.Children[0].MoveDownRight(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveDownRight did an invalid move at puzzle index 3")
	}

	//  ? | ? | ? | ?
	//	? | ? | ? | 0
	//	? | ? | ? | ?
	p.Children[0].Value = 7
	p.Children[0].MoveDownRight(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveDownRight did an invalid move at puzzle index 7")
	}

	//  ? | ? | ? | ?
	//	? | ? | ? | ?
	//	? | ? | 0 | ?
	p.Children[0].Value = 10
	p.Children[0].MoveDownRight(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveDownRight did an invalid move at puzzle index 10")
	}

	//  ? | ? | ? | ?
	//	? | ? | ? | ?
	//	? | 0 | ? | ?
	p.Children[0].Value = 9
	p.Children[0].MoveDownRight(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveDownRight did an invalid move at puzzle index 9")
	}

	//  ? | ? | ? | ?
	//	? | ? | ? | ?
	//	0 | ? | ? | ?
	p.Children[0].Value = 8
	p.Children[0].MoveDownRight(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveDownRight did an invalid move at puzzle index 8")
	}
}

func TestMovedDown(t *testing.T) {
	i := []int{1, 2, 3, 4, 5, 6, 0, 8, 9, 10, 7, 11}
	p := NewPuzzle(i)

	//  1 | 2 | 3 | 4
	//	5 | 6 | 0 | 8
	//	9 | 10| 7 | 11

	// find the position of the "0"
	// this is done in the GenerateMoves()
	for i := 0; i < len(p.Puzzle); i++ {
		if p.Puzzle[i] == 0 {
			p.Value = i
		}
	}

	p.MoveDown(p.Puzzle, p.Value)
	up := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 0, 11}
	if AreTheSame(p.Children[0].Puzzle, up) != true {
		t.Errorf("MoveDown did not move the puzzle correctly")
	}

	// TEST INVALID MOVES

	// p.Children[0].Puzzle looks like this:
	//  1 | 2 | 3 | 4
	//	5 | 6 | 7 | 8
	//	9 | 10| 0 | 11
	// let's try to make an invalid move, and
	// move it up
	p.Children[0].Value = 10
	p.Children[0].MoveDown(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveDown did an invalid move at puzzle index 10")
	}

	//  ? | ? | ? | ?
	//	? | ? | ? | ?
	//	? | ? | ? | 0
	p.Children[0].Value = 11
	p.Children[0].MoveDown(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveDown did an invalid move at puzzle index 11")
	}

	//  ? | ? | ? | ?
	//	? | ? | ? | ?
	//	? | 0 | ? | ?
	p.Children[0].Value = 9
	p.Children[0].MoveDown(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveDown did an invalid move at puzzle index 9")
	}

	//  ? | ? | ? | ?
	//	? | ? | ? | ?
	//	0 | ? | ? | ?
	p.Children[0].Value = 8
	p.Children[0].MoveDown(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveDown did an invalid move at puzzle index 8")
	}
}

func TestMovedDownLeft(t *testing.T) {
	i := []int{1, 2, 3, 4, 5, 6, 0, 8, 9, 10, 7, 11}
	p := NewPuzzle(i)

	//  1 | 2 | 3 | 4
	//	5 | 6 | 0 | 8
	//	9 | 10| 7 | 11

	// find the position of the "0"
	// this is done in the GenerateMoves()
	for i := 0; i < len(p.Puzzle); i++ {
		if p.Puzzle[i] == 0 {
			p.Value = i
		}
	}

	p.MoveDownLeft(p.Puzzle, p.Value)
	up := []int{1, 2, 3, 4, 5, 6, 10, 8, 9, 0, 7, 11}
	if AreTheSame(p.Children[0].Puzzle, up) != true {
		t.Errorf("MoveDownLeft did not move the puzzle correctly")
	}

	// TEST INVALID MOVES

	// p.Children[0].Puzzle looks like this:
	//  1 | 2 | 3 | 4
	//	5 | 6 | 10| 8
	//	9 | 0 | 7 | 11
	// let's try to make an invalid move, and
	// move it up
	p.Children[0].Value = 9
	p.Children[0].MoveDownLeft(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveDownLeft did an invalid move at puzzle index 9")
	}

	//  ? | ? | ? | ?
	//	? | ? | ? | ?
	//	? | ? | ? | 0
	p.Children[0].Value = 11
	p.Children[0].MoveDownLeft(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveDownLeft did an invalid move at puzzle index 11")
	}

	//  ? | ? | ? | ?
	//	? | ? | ? | ?
	//	0 | ? | ? | ?
	p.Children[0].Value = 8
	p.Children[0].MoveDownLeft(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveDownLeft did an invalid move at puzzle index 8")
	}

	//  0 | ? | ? | ?
	//	? | ? | ? | ?
	//	? | ? | ? | ?
	p.Children[0].Value = 0
	p.Children[0].MoveDownLeft(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveDownLeft did an invalid move at puzzle index 0")
	}

	//  ? | ? | ? | ?
	//	0 | ? | ? | ?
	//	? | ? | ? | ?
	p.Children[0].Value = 4
	p.Children[0].MoveDownLeft(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveDownLeft did an invalid move at puzzle index 4")
	}
}

func TestMovedLeft(t *testing.T) {
	i := []int{1, 2, 3, 4, 5, 6, 0, 8, 9, 10, 7, 11}
	p := NewPuzzle(i)

	//  1 | 2 | 3 | 4
	//	5 | 6 | 0 | 8
	//	9 | 10| 7 | 11

	// find the position of the "0"
	// this is done in the GenerateMoves()
	for i := 0; i < len(p.Puzzle); i++ {
		if p.Puzzle[i] == 0 {
			p.Value = i
		}
	}

	p.MoveLeft(p.Puzzle, p.Value)
	up := []int{1, 2, 3, 4, 5, 0, 6, 8, 9, 10, 7, 11}
	if AreTheSame(p.Children[0].Puzzle, up) != true {
		t.Errorf("MoveLeft did not move the puzzle correctly")
	}

	// TEST INVALID MOVES

	//  ? | ? | ? | ?
	//	? | ? | ? | ?
	//	0 | ? | ? | ?
	p.Children[0].Value = 8
	p.Children[0].MoveLeft(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveLeft did an invalid move at puzzle index 8")
	}

	//  0 | ? | ? | ?
	//	? | ? | ? | ?
	//	? | ? | ? | ?
	p.Children[0].Value = 0
	p.Children[0].MoveLeft(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveLeft did an invalid move at puzzle index 0")
	}

	//  ? | ? | ? | ?
	//	0 | ? | ? | ?
	//	? | ? | ? | ?
	p.Children[0].Value = 4
	p.Children[0].MoveLeft(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveLeft did an invalid move at puzzle index 4")
	}
}

func TestMovedUpLeft(t *testing.T) {
	i := []int{1, 2, 3, 4, 5, 6, 0, 8, 9, 10, 7, 11}
	p := NewPuzzle(i)

	//  1 | 2 | 3 | 4
	//	5 | 6 | 0 | 8
	//	9 | 10| 7 | 11

	// find the position of the "0"
	// this is done in the GenerateMoves()
	for i := 0; i < len(p.Puzzle); i++ {
		if p.Puzzle[i] == 0 {
			p.Value = i
		}
	}

	p.MoveUpLeft(p.Puzzle, p.Value)
	up := []int{1, 0, 3, 4, 5, 6, 2, 8, 9, 10, 7, 11}
	if AreTheSame(p.Children[0].Puzzle, up) != true {
		t.Errorf("MoveUpLeft did not move the puzzle correctly")
	}

	// TEST INVALID MOVES

	//  ? | ? | ? | ?
	//	? | ? | ? | ?
	//	0 | ? | ? | ?
	p.Children[0].Value = 8
	p.Children[0].MoveLeft(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveUpLeft did an invalid move at puzzle index 8")
	}

	//  0 | ? | ? | ?
	//	? | ? | ? | ?
	//	? | ? | ? | ?
	p.Children[0].Value = 0
	p.Children[0].MoveUpLeft(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveUpLeft did an invalid move at puzzle index 0")
	}

	//  ? | ? | ? | ?
	//	0 | ? | ? | ?
	//	? | ? | ? | ?
	p.Children[0].Value = 4
	p.Children[0].MoveUpLeft(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveUpLeft did an invalid move at puzzle index 4")
	}

	//  ? | ? | ? | 0
	//	? | ? | ? | ?
	//	? | ? | ? | ?
	p.Children[0].Value = 3
	p.Children[0].MoveUpLeft(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveUpLeft did an invalid move at puzzle index 3")
	}

	//  ? | ? | 0 | ?
	//	? | ? | ? | ?
	//	? | ? | ? | ?
	p.Children[0].Value = 2
	p.Children[0].MoveUpLeft(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveUpLeft did an invalid move at puzzle index 3")
	}

	//  ? | 0 | ? | ?
	//	? | ? | ? | ?
	//	? | ? | ? | ?
	p.Children[0].Value = 1
	p.Children[0].MoveUpLeft(p.Children[0].Puzzle, p.Children[0].Value)
	if len(p.Children[0].Children) != 0 {
		t.Errorf("MoveUpLeft did an invalid move at puzzle index 3")
	}
}

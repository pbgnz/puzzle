package puzzle

import "testing"

func BenchmarkDFS(b *testing.B) {
	i := []int{5, 1, 2, 3, 9, 6, 7, 4, 0, 10, 11, 8}
	p := NewPuzzle(i)
	DepthFirstSearch(p)
}

func BenchmarkBFSH1(b *testing.B) {
	i := []int{5, 1, 2, 3, 9, 6, 7, 4, 0, 10, 11, 8}
	p := NewPuzzle(i)
	BestFirstSearch(p, 1)
}

// func BenchmarkBFSH2(b *testing.B) {
// 	i := []int{5, 1, 2, 3, 9, 6, 7, 4, 0, 10, 11, 8}
// 	p := NewPuzzle(i)
// 	BestFirstSearch(p, 2)
// }

func BenchmarkASH1(b *testing.B) {
	i := []int{5, 1, 2, 3, 9, 6, 7, 4, 0, 10, 11, 8}
	p := NewPuzzle(i)
	As(p, 1)
}

func BenchmarkASH2(b *testing.B) {
	i := []int{5, 1, 2, 3, 9, 6, 7, 4, 0, 10, 11, 8}
	p := NewPuzzle(i)
	As(p, 2)
}

func TestDepthFirstSearch(t *testing.T) {

	//  1 | 2 | 3 | 0
	//	5 | 6 | 7 | 4
	//	9 | 10| 11| 8

	// DFS algorithm should finish the above puzzle
	// in two moves (move 0 DOWN -> move 0 DOWN)
	i := []int{1, 2, 3, 0, 5, 6, 7, 4, 9, 10, 11, 8}
	p := NewPuzzle(i)

	// the results should be the solution path;
	// [move2 (goal state), move1, intial state]
	r := DepthFirstSearch(p)

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

func TestDepthFirstSearchComplex(t *testing.T) {

	//  5 | 1 | 2 | 3
	//	9 | 6 | 7 | 4
	//	0 | 10| 11| 8
	// DFS algorithm should finish the above puzzle
	// in 7 moves: UP -> UP -> RIGHT -> RIGHT -> RIGHT -> DOWN - DOWN

	i := []int{5, 1, 2, 3, 9, 6, 7, 4, 0, 10, 11, 8}
	p := NewPuzzle(i)

	r := DepthFirstSearch(p)

	// results should return 7 moves + initial state
	if len(r) != 8 {
		t.Errorf("Expected the solution path to be size 8, but was size %v", len(r))
	}

	move1 := []int{5, 1, 2, 3, 0, 6, 7, 4, 9, 10, 11, 8}
	move2 := []int{0, 1, 2, 3, 5, 6, 7, 4, 9, 10, 11, 8}
	move3 := []int{1, 0, 2, 3, 5, 6, 7, 4, 9, 10, 11, 8}
	move4 := []int{1, 2, 0, 3, 5, 6, 7, 4, 9, 10, 11, 8}
	move5 := []int{1, 2, 3, 0, 5, 6, 7, 4, 9, 10, 11, 8}
	move6 := []int{1, 2, 3, 4, 5, 6, 7, 0, 9, 10, 11, 8}
	move7 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 0}

	// [move7, move6, move5, move4, move3, move2, move1, intial state]
	if AreTheSame(r[0].Puzzle, move7) != true {
		t.Errorf("Expected the solution path to return the 7th move")
	}

	if AreTheSame(r[1].Puzzle, move6) != true {
		t.Errorf("Expected the solution path to return the 6th move")
	}

	if AreTheSame(r[2].Puzzle, move5) != true {
		t.Errorf("Expected the solution path to return the 5th move")
	}

	if AreTheSame(r[3].Puzzle, move4) != true {
		t.Errorf("Expected the solution path to return the 4th move")
	}

	if AreTheSame(r[4].Puzzle, move3) != true {
		t.Errorf("Expected the solution path to return the 3rd move")
	}

	if AreTheSame(r[5].Puzzle, move2) != true {
		t.Errorf("Expected the solution path to return the 2nd move")
	}

	if AreTheSame(r[6].Puzzle, move1) != true {
		t.Errorf("Expected the solution path to return the 1st move")
	}

	if AreTheSame(r[7].Puzzle, i) != true {
		t.Errorf("Expected the solution path to return the intial state")
	}
}

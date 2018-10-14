package puzzle

import "testing"

func TestHeuristic1(t *testing.T) {
	// Heuristic1 is the hamming distance
	// that is the number of misplaced tiles

	// 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 = 12
	i := []int{5, 3, 7, 0, 11, 1, 6, 10, 4, 9, 2, 8}
	h := Heuristic1(i)

	if h != 12 {
		t.Errorf("Expected a hamming distance of 12, but got %v", h)
	}

	// 0 + 0 + 0 + 0 + 0 + 0 + 0 + 0 + 0 + 0 + 1 + 1 = 2
	i2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 0, 11}
	h2 := Heuristic1(i2)

	if h2 != 2 {
		t.Errorf("Expected a hamming distance of 2, but got %v", h2)
	}

	// 0 + 0 + 0 + 0 + 0 + 0 + 0 + 0 + 0 + 0 + 0 + 0 = 0
	i3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 0}
	h3 := Heuristic1(i3)

	if h3 != 0 {
		t.Errorf("Expected a hamming distance of 0, but got %v", h3)
	}
}

func TestHeuristic2(t *testing.T) {
	// Heuristic2 is the sum of permutation inversions + hamming distance

	// 4 + 2 + 4 + - + 7 + 0 + 2 + 4 + 1 + 2 + 0 + 0 = 26
	// 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 = 12
	i := []int{5, 3, 7, 0, 11, 1, 6, 10, 4, 9, 2, 8}
	h := Heuristic2(i)

	if h != 38 {
		t.Errorf("Expected a sum of permutation inversions distance of 38, but got %v", h)
	}

	// 1 + 0 + 0 + 0 + 0 + 0 + 0 + 0 + 0 + 0 + 0 + - = 1
	// 1 + 1 + 0 + 0 + 0 + 0 + 0 + 0 + 0 + 0 + 0 + 0 = 2
	i2 := []int{2, 1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 0}
	h2 := Heuristic2(i2)

	if h2 != 3 {
		t.Errorf("Expected a sum of permutation inversions distance of 3, but got %v", h2)
	}

	// 0 + 0 + 0 + 0 + 0 + 0 + 0 + 0 + 0 + 0 + 0 + - = 0
	// 0 + 0 + 0 + 0 + 0 + 0 + 0 + 0 + 0 + 0 + 0 + 0 = 2
	i3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 0}
	h3 := Heuristic2(i3)

	if h3 != 0 {
		t.Errorf("Expected a sum of permutation inversions of 0, but got %v", h3)
	}
}

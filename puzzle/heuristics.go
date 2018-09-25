package puzzle

var goalState = map[int]int{
	0:  1,
	1:  2,
	2:  3,
	3:  4,
	4:  5,
	5:  6,
	6:  7,
	7:  8,
	8:  9,
	9:  10,
	10: 11,
	11: 0,
}

// Heuristic1 is the number of misplaced tiles
func Heuristic1(p []int) int {
	m := 0
	for i := 0; i < len(p)-1; i++ {
		if p[i] == goalState[i] {
			continue
		}
		m++
	}
	return m
}

// Heuristic2 sum of permutation inversions
func Heuristic2(p []int) int {
	m := 0
	for i := 0; i < len(p)-1; i++ {
		if p[i] == goalState[i] {
			continue
		}
	}
	return m
}

package puzzle

// TODO: f(N) must be admissible

// Heuristic1 is the number of misplaced tiles
func Heuristic1(p []int) int {
	m := 0
	for i := 0; i < len(p); i++ {
		if p[i] != goalState[i] {
			m++
		}
	}
	return m
}

// Heuristic2 sum of permutation inversions
func Heuristic2(p []int) int {
	m := 0
	for i := 0; i < len(p); i++ {
		if p[i] != 0 {
			for j := i + 1; j < len(p); j++ {
				if p[j] == 0 {
					continue
				}
				if inverseGoalState[p[j]] < inverseGoalState[p[i]] {
					m++
				}
			}
		}
	}
	m = +Heuristic1(p);
	return m
}

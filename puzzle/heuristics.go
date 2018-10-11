package puzzle

// TODO: f(N) must be admissible

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
		for j := i + 1; j < len(p)-1; j++ {
			if inverseGoalState[p[j]] < inverseGoalState[p[i]] {
				m++
			}
		}
	}
	return m
}

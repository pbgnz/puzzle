package puzzle

const (
	// NumberRows of the board
	NumberRows = 3
	// NumberColumns of the board
	NumberColumns = 4
)

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

var inverseGoalState = map[int]int{
	1:  0,
	2:  1,
	3:  2,
	4:  3,
	5:  4,
	6:  5,
	7:  6,
	8:  7,
	9:  8,
	10: 9,
	11: 10,
	0:  11,
}

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

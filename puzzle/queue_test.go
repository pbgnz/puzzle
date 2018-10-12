package puzzle

import (
	"container/heap"
	"reflect"
	"testing"
)

func TestPop(t *testing.T) {
	// create a pq with two items inside
	pq := make(PriorityQueue, 2)

	// create two puzzles
	i := []int{5, 3, 7, 0, 11, 1, 6, 10, 4, 9, 2, 8}
	p := NewPuzzle(i)
	p.Heuristic = 10
	item := &Item{
		Node:      p,
		Heuristic: p.Heuristic,
	}

	p2 := NewPuzzle(i)
	p2.Heuristic = 12
	item2 := &Item{
		Node:      p2,
		Heuristic: p2.Heuristic,
	}

	pq[0] = item
	pq[1] = item2
	heap.Init(&pq)

	// pop the items and verify if lowest heuristic poped first
	pop1 := heap.Pop(&pq).(*Item)
	if reflect.DeepEqual(item, pop1) != true {
		t.Errorf("Heuristic queue did not pop the lowest Heuristic element")
	}

	pop2 := heap.Pop(&pq).(*Item)
	if reflect.DeepEqual(item2, pop2) != true {
		t.Errorf("Heuristic queue did not pop the lowest Heuristic element")
	}
}

func TestPush(t *testing.T) {
	// create an empty pq
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// create two puzzles
	i := []int{5, 3, 7, 0, 11, 1, 6, 10, 4, 9, 2, 8}
	p := NewPuzzle(i)
	p.Heuristic = 10
	item := &Item{
		Node:      p,
		Heuristic: p.Heuristic,
	}

	p2 := NewPuzzle(i)
	p2.Heuristic = 12
	item2 := &Item{
		Node:      p2,
		Heuristic: p2.Heuristic,
	}

	// push the two items
	heap.Push(&pq, item)
	if pq.Len() != 1 {
		t.Errorf("The length of the pq did not update properly; expected 1 but got %v", pq.Len())
	}
	heap.Push(&pq, item2)
	if pq.Len() != 2 {
		t.Errorf("The length of the pq did not update properly; expected 2 but got %v", pq.Len())
	}

	pop1 := heap.Pop(&pq).(*Item)
	if reflect.DeepEqual(item, pop1) != true {
		t.Errorf("Heuristic queue did not pop the lowest Heuristic element")
	}

	pop2 := heap.Pop(&pq).(*Item)
	if reflect.DeepEqual(item2, pop2) != true {
		t.Errorf("Heuristic queue did not pop the lowest Heuristic element")
	}
}

func TestUpdate(t *testing.T) {
	// create a pq with two items inside
	pq := make(PriorityQueue, 2)

	// create two puzzles
	i := []int{5, 3, 7, 0, 11, 1, 6, 10, 4, 9, 2, 8}
	p := NewPuzzle(i)
	p.Heuristic = 10

	p2 := NewPuzzle(i)
	p2.Heuristic = 12

	item := &Item{
		Node:      p,
		Heuristic: p.Heuristic,
	}

	item2 := &Item{
		Node:      p2,
		Heuristic: p2.Heuristic,
	}

	pq[0] = item
	pq[1] = item2
	heap.Init(&pq)

	// update the priorities of the two items inside the pq
	pq.update(item, p, 100)
	pq.update(item2, p2, 1)

	// verify the items were poped in the right order
	pop1 := heap.Pop(&pq).(*Item)
	if reflect.DeepEqual(item2, pop1) != true {
		t.Errorf("Heuristic queue did not pop the lowest Heuristic element")
	}

	pop2 := heap.Pop(&pq).(*Item)
	if reflect.DeepEqual(item, pop2) != true {
		t.Errorf("Heuristic queue did not pop the lowest Heuristic element")
	}
}

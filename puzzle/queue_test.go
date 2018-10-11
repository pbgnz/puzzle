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

	p2 := NewPuzzle(i)
	p2.Heuristic = 12

	item := &Item{
		Value:    p,
		Priority: p.Heuristic,
	}

	item2 := &Item{
		Value:    p2,
		Priority: p2.Heuristic,
	}

	pq[0] = item
	pq[1] = item2
	heap.Init(&pq)

	// pop the items and verify if lowest heuristic poped first
	pop1 := pq.Pop().(*Item)
	if reflect.DeepEqual(item, pop1) != true {
		t.Errorf("Priority queue did not pop the lowest priority element")
	}

	pop2 := pq.Pop().(*Item)
	if reflect.DeepEqual(item2, pop2) != true {
		t.Errorf("Priority queue did not pop the lowest priority element")
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

	p2 := NewPuzzle(i)
	p2.Heuristic = 12

	item := &Item{
		Value:    p,
		Priority: p.Heuristic,
	}

	item2 := &Item{
		Value:    p2,
		Priority: p2.Heuristic,
	}

	// push the two items
	pq.Push(item)
	heap.Push(&pq, item)
	pq.Push(item2)
	heap.Push(&pq, item2)

	// verify the items were poped in the right order
	pop1 := pq.Pop().(*Item)
	if reflect.DeepEqual(item, pop1) != true {
		t.Errorf("Priority queue did not pop the lowest priority element")
	}

	pop2 := pq.Pop().(*Item)
	if reflect.DeepEqual(item2, pop2) != true {
		t.Errorf("Priority queue did not pop the lowest priority element")
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
		Value:    p,
		Priority: p.Heuristic,
	}

	item2 := &Item{
		Value:    p2,
		Priority: p2.Heuristic,
	}

	pq[0] = item
	pq[1] = item2
	heap.Init(&pq)

	// update the priorities of the two items inside the pq
	pq.update(item, p, 100)
	pq.update(item2, p2, 1)

	// verify the items were poped in the right order
	pop1 := pq.Pop().(*Item)
	if reflect.DeepEqual(item2, pop1) != true {
		t.Errorf("Priority queue did not pop the lowest priority element")
	}

	pop2 := pq.Pop().(*Item)
	if reflect.DeepEqual(item, pop2) != true {
		t.Errorf("Priority queue did not pop the lowest priority element")
	}
}

package puzzle

import (
	"container/heap"
	"reflect"
)

// An Item is something we manage in a priority queue.
type Item struct {
	Node      *Node // The value of the item; arbitrary.
	Heuristic int   // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	Index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Heuristic < pq[j].Heuristic
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// Push an element to the pq
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

// Pop an element from the pq
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value *Node, priority int) {
	item.Node = value
	item.Heuristic = priority
	heap.Fix(pq, item.Index)
}

// Contains checks if it is there
func (pq PriorityQueue) Contains(n *Node) (bool, *Item) {
	for i := 0; i < len(pq); i++ {
		if AreTheSame(pq[i].Node.Puzzle, n.Puzzle) {
			if reflect.DeepEqual(pq[i].Node, n) {
				return true, pq[i]
			}
		}
	}
	return false, nil
}

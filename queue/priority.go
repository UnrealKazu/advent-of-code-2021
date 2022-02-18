package queue

import "container/heap"

type Item struct {
	I, J     int
	Priority int
	Index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)

	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	item.Index = -1

	*pq = old[:n-1]

	return item
}

func (pq *PriorityQueue) Push(i interface{}) {
	n := len(*pq)
	item := i.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Update(item *Item, prio int) {
	item.Priority = prio
	heap.Fix(pq, item.Index)
}

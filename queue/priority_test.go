package queue

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	items := []*Item{
		{Priority: 5},
		{Priority: 1},
		{Priority: 100},
		{Priority: 50},
		{Priority: 26},
	}

	pq := make(PriorityQueue, len(items))

	for i, it := range items {
		pq[i] = it
	}

	heap.Init(&pq)

	heap.Push(&pq, &Item{
		Priority: 44,
	})

	for _, i := range pq {
		fmt.Println(i)
	}
}

// func TestAnother(t *testing.T) {
// 	items := []*AnotherThing{
// 		{Name: "wablaat", Item: Item{
// 			Priority: 50,
// 		}},
// 		{Name: "je moeder", Item: Item{
// 			Priority: 24,
// 		}},
// 		{Name: "je vader", Item: Item{
// 			Priority: 70,
// 		}},
// 	}

// 	pq := make(PriorityQueue, len(items))

// 	for i, it := range items {
// 		pq[i] = &it.Item
// 	}

// 	heap.Init(&pq)

// 	heap.Push(&pq, &Item{
// 		Priority: 44,
// 	})

// 	for _, i := range pq {
// 		fmt.Println(i)
// 	}
// }

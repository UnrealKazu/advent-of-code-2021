// Package graph provides a graph struct with a lowest cost traversal algorithm (Dijkstra) on it
package graph

import (
	"fmt"
	"math"
	"sort"
)

const (
	Debug = false
)

type Node struct {
	I, J, Distance int
}

type Graph struct {
	Matrix   [][]int
	Distance [][]int
	Queue    []*Node
}

func New(in [][]int) *Graph {
	return &Graph{
		Matrix: in,
		Queue:  []*Node{},
	}
}

func NewNode(i, j int) *Node {
	return &Node{
		I: i,
		J: j,
	}
}

// GetLowestRiskPath traverses the matrix in the graph using Dijkstra's. It determines the lowest cost path and returns it.
func (t *Graph) GetLowestRiskPath() int {
	t.Distance = [][]int{}

	for i := 0; i < len(t.Matrix); i++ {
		t.Distance = append(t.Distance, []int{})
		for j := 0; j < len(t.Matrix[i]); j++ {
			t.Distance[i] = append(t.Distance[i], math.MaxInt)
		}
	}

	// push the source (i.e. 0,0) to the queue
	t.Queue = append(t.Queue, NewNode(0, 0))
	t.Distance[0][0] = 0

	for len(t.Queue) > 0 {
		cur := t.Queue[0]

		// remove the current node from the queue
		t.Queue = t.Queue[1:]

		// check left
		if cur.J > 0 && t.Distance[cur.I][cur.J]+t.Matrix[cur.I][cur.J-1] < t.Distance[cur.I][cur.J-1] {
			t.Distance[cur.I][cur.J-1] = t.Distance[cur.I][cur.J] + t.Matrix[cur.I][cur.J-1]
			t.Queue = append(t.Queue, NewNode(cur.I, cur.J-1))
		}

		// check top
		if cur.I > 0 && t.Distance[cur.I][cur.J]+t.Matrix[cur.I-1][cur.J] < t.Distance[cur.I-1][cur.J] {
			t.Distance[cur.I-1][cur.J] = t.Distance[cur.I][cur.J] + t.Matrix[cur.I-1][cur.J]
			t.Queue = append(t.Queue, NewNode(cur.I-1, cur.J))
		}

		// check right
		if cur.J < len(t.Matrix[cur.I])-1 && t.Distance[cur.I][cur.J]+t.Matrix[cur.I][cur.J+1] < t.Distance[cur.I][cur.J+1] {
			t.Distance[cur.I][cur.J+1] = t.Distance[cur.I][cur.J] + t.Matrix[cur.I][cur.J+1]
			t.Queue = append(t.Queue, NewNode(cur.I, cur.J+1))
		}

		// check down
		if cur.I < len(t.Matrix)-1 && t.Distance[cur.I][cur.J]+t.Matrix[cur.I+1][cur.J] < t.Distance[cur.I+1][cur.J] {
			t.Distance[cur.I+1][cur.J] = t.Distance[cur.I][cur.J] + t.Matrix[cur.I+1][cur.J]
			t.Queue = append(t.Queue, NewNode(cur.I+1, cur.J))
		}

		t.sortQueue()
	}

	if Debug {
		// print the final distance matrix
		for i := 0; i < len(t.Distance); i++ {
			for j := 0; j < len(t.Distance[i]); j++ {
				fmt.Printf("%d,", t.Distance[i][j])
			}
			fmt.Println()
		}
	}

	// return the distance at the bottom right cell
	return t.Distance[len(t.Distance)-1][len(t.Distance[len(t.Distance)-1])-1]
}

func (t *Graph) sortQueue() {
	sort.SliceStable(t.Queue, func(i, j int) bool {
		ni := t.Queue[i]
		nj := t.Queue[j]
		return t.Distance[ni.I][ni.J] < t.Distance[nj.I][nj.J]
	})
}

package main

import (
	"fmt"
	"strings"

	"github.com/advent-of-code-2021/day12/first/graph"
	"github.com/advent-of-code-2021/input"
)

func main() {
	in := input.Read("input.txt")

	nrofPaths := GetNrofSmallCavePaths(in)

	fmt.Printf("Number of paths that only visit small caves at most once %d\n", nrofPaths)
}

// GetNrofSmallCavePaths builds a graph on the input set, and returns a set of distinct paths
// that does not cover small caves multiple times
func GetNrofSmallCavePaths(in []string) int {
	// build up the graph, create the nodes, and link
	// all nodes together
	g := graph.New()

	for _, r := range in {
		spl := strings.Split(r, "-")
		sn1, sn2 := spl[0], spl[1]

		n1 := g.AddNode(sn1)
		n2 := g.AddNode(sn2)

		g.AddRelation(n1, n2)
	}

	return g.GetDistinctPaths(g.Nodes["start"], []string{})
}

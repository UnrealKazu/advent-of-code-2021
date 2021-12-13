package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/advent-of-code-2021/day12/second/graph"
	"github.com/advent-of-code-2021/input"
)

// Note: this part 2 solution is incredibly ugly, but I was strapped for time and had to sort of
// brute force the "one small cave more than once" solution by allowing for duplicate paths and
// filtering them out later. Very inefficient, but I will clean it up later.
func main() {
	in := input.Read("input.txt")

	numPaths := GetNrofSmallCavePaths(in)

	fmt.Printf("Number of paths %d\n", numPaths)
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

	// get all paths
	paths := g.GetPaths("start,", g.Nodes["start"], []string{}, true)

	// split them so that we can do a unique check
	splPaths := strings.Split(paths, "\n")

	sort.Strings(splPaths)

	unique := []string{splPaths[0]}

	for i := 1; i < len(splPaths); i++ {
		if splPaths[i] != splPaths[i-1] {
			unique = append(unique, splPaths[i])
		}
	}

	// remove the first entry because it's always empty
	return len(unique) - 1
}

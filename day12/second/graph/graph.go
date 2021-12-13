// Package graph provides graph traversal functions
package graph

import "unicode"

type Node struct {
	Label string
	Out   map[string]*Node
	In    map[string]*Node
	Large bool
}

type Graph struct {
	Nodes map[string]*Node
}

func New() *Graph {
	return &Graph{
		Nodes: map[string]*Node{},
	}
}

func (g *Graph) GetPaths(path string, n *Node, vis []string, secPossible bool) string {
	if n.Label == "end" {
		// we've reached the end, so this is one full path
		return path + "\n"
	}

	newPath := ""
	for _, nn := range n.Out {
		contains := false
		for _, c := range vis {
			if c == nn.Label {
				contains = true
				break
			}
		}

		if !contains {
			nvis := make([]string, len(vis))
			copy(nvis, vis)
			if !nn.Large {
				// only if the node is a small node do we want to track it
				if secPossible {
					// we have two options here:
					// 1. Add this node to the visited list, and keep the possibility to visit another node twice
					// 2. Do not add this node to the visited list, but remove the possibility to visit another node twice

					// 2
					newPath += g.GetPaths(path+nn.Label+",", nn, nvis, false)

					nvis = append(vis, nn.Label)
					newPath += g.GetPaths(path+nn.Label+",", nn, nvis, true)
				} else {
					// a second visit is no longer possible, so add the node to visited
					nvis = append(vis, nn.Label)
					newPath += g.GetPaths(path+nn.Label+",", nn, nvis, secPossible)
				}
			} else {
				newPath += g.GetPaths(path+nn.Label+",", nn, nvis, secPossible)
			}
		}
	}

	return newPath
}

func (g *Graph) AddRelation(n1, n2 *Node) {
	if n2.Label == "start" || n1.Label == "end" {
		// direction is important for these two special cases
		n1, n2 = n2, n1
	}

	if n1.Label == "start" || n2.Label == "end" {
		// only outgoing from n1 to n2
		n1.Out[n2.Label] = n2
		n2.In[n1.Label] = n1
	} else {
		// bi-directional
		n1.Out[n2.Label] = n2
		n2.Out[n1.Label] = n1
		n2.In[n1.Label] = n1
		n1.In[n2.Label] = n1
	}
}

// AddNode adds a node if it does not already exist in the graph
func (g *Graph) AddNode(label string) *Node {
	if n, ok := g.Nodes[label]; !ok {
		// node does not exist yet, add it
		var large bool
		if unicode.IsUpper(rune(label[0])) {
			// check if the first rune of the label is uppercase
			// if so, this is a large node
			large = true
		} else {
			large = false
		}

		nn := Node{
			Label: label,
			Out:   make(map[string]*Node, 0),
			In:    make(map[string]*Node, 0),
			Large: large,
		}

		g.Nodes[label] = &nn

		return &nn
	} else {
		return n
	}
}

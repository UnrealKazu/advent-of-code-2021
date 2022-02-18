// Package tree provides a binary tree for performing snailfish mathematical additions on
package tree

import (
	"fmt"
	"strconv"
)

type Node struct {
	Key    int
	Value  int
	Length int
	Parent *Node
	Left   *Node
	Right  *Node
}

type Tree struct {
	nrofAdded int
	Root      *Node
}

func New(l, r int) *Tree {
	pn := &Node{
		Key:    0,
		Length: 1,
	}

	ln := &Node{
		Key:    1,
		Value:  l,
		Length: 0,
		Parent: pn,
	}

	rn := &Node{
		Key:    2,
		Value:  r,
		Length: 0,
		Parent: pn,
	}

	pn.Left = ln
	pn.Right = rn

	return &Tree{
		Root:      pn,
		nrofAdded: 3,
	}
}

func Generate(s string) *Tree {
	q := []*Node{}
	openNode := false
	for _, ss := range s {
		switch string(ss) {
		case "[":
			// new node
			q = append(q, &Node{
				Key:    0,
				Length: 0,
			})

			// also open node mode if any is open
			openNode = false
		case ",":
			// comma, don't do anything
		case "]":
			// close node, and pop it from the queue
			openNode = false

			if len(q) <= 1 {
				// this is the last node, so noop
				break
			}

			first := q[len(q)-1]
			second := q[len(q)-2]

			first.Parent = second

			if second.Left == nil {
				second.Left = first
			} else if second.Right == nil {
				second.Right = first
			} else {
				panic("Unexpected situation. Exiting.")
			}

			// pop the latest node
			q = q[:len(q)-2]
		default:
			num, _ := strconv.Atoi(string(ss))
			// must be a number
			if !openNode {
				// not in node mode, so this is a leftmost pair
				l := &Node{
					Key:    0,
					Length: 0,
					Value:  num,
				}

				// add the leftmost pair to the node on the top of the stack
				top := q[len(q)-1]
				top.Left = l
				top.Length++

				openNode = true
			} else {
				// we're in node mode, so this is the right hand of a pair
				r := &Node{
					Key:    0,
					Length: 0,
					Value:  num,
				}

				top := q[len(q)-1]
				top.Right = r
			}
		}
	}

	t := &Tree{
		Root: q[0],
	}

	return t
}

func (t *Tree) Add(l, r int) {
	// first, create a new node for the new pair
	pn := &Node{
		Key:    t.nrofAdded + 1,
		Length: 1,
	}

	ln := &Node{
		Key:    t.nrofAdded + 2,
		Value:  l,
		Length: 0,
		Parent: pn,
	}

	rn := &Node{
		Key:    t.nrofAdded + 3,
		Value:  r,
		Length: 0,
		Parent: pn,
	}

	pn.Left = ln
	pn.Right = rn

	// replace the root with a new root
	// with old root as the left node,
	// and the addition as the right node
	nr := &Node{
		Key:    t.nrofAdded + 4,
		Length: t.Root.Length + 1,
		Left:   t.Root,
		Right:  pn,
	}

	nr.Left.Parent = nr
	nr.Right.Parent = nr

	t.Root = nr
	t.nrofAdded += 4
}

func (t *Tree) Reduce() {
	t.reduceNode(t.Root)
}

func (t *Tree) reduceNode(n *Node) {
	if n.Length >= 4 {
		// we need to explode this pair
		t.explode(n)
	}

	if isLeaf(n) {
		if n.Value >= 10 {
			// we need to split this pair
			t.split(n)
		}

		// further reduce if necessary
		if n.Left != nil {
			t.reduceNode(n.Left)
		}

		// further reduce if necessary
		if n.Right != nil {
			t.reduceNode(n.Right)
		}
	}

}

func (t *Tree) explode(n *Node) {
	// sanity check to make sure we're operating on a valid node
	if !isLeaf(n.Left) && !isLeaf(n.Right) {
		fmt.Printf("Leaf operation on non-leaf %d\n", n.Key)
		panic("Invalid operation. Tree is corrupt. Exiting")
	}

	// add the left value to the first regular number to the left of this node
	// (i.e. the left of the parent, if it's not this node)
	if n.Parent.Left != nil && n.Parent.Left.Key != n.Key {
		// make sure that we do not add the value to ourselves
		n.Parent.Left.Value += n.Left.Value

		// TODO: check if the modified node needs another operation
	}

	// add the right value to the first regular number to the right of this node
	if n.Parent.Right != nil && n.Parent.Right.Key != n.Key {
		n.Parent.Right.Value += n.Right.Value

		// TODO: check if the modified node needs another operation
	}

	// next, the current node no longer has leafs and becomes a leaf
	n.Left = nil
	n.Right = nil
	n.Value = 0
}

func (t *Tree) split(n *Node) {

}

func isLeaf(n *Node) bool {
	return n.Left == nil && n.Right == nil
}

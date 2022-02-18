package tree

import "testing"

func TestGenerate_WithSmallValidInput_ShouldGenerateCorrectTree(t *testing.T) {
	in := "[2,9]"

	tt := Generate(in)

	if tt.Root.Left.Value != 2 {
		t.Errorf("Unexpected. Expected %d, got %d", 2, tt.Root.Left.Value)
	}
}

func TestGenerate_WithLargeValidInput_ShouldGenerateCorrectTree(t *testing.T) {
	in := "[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]"

	tt := Generate(in)

	if tt.Root.Left.Value != 7 {
		t.Errorf("Unexpected. Expected %d, got %d", 2, tt.Root.Left.Value)
	}
}

func TestNew_WithFreshRoot_ShouldCreateValidTreeLeafLeft(t *testing.T) {
	tr := New(5, 6)

	exp := 5
	act := tr.Root.Left.Value

	if act != exp {
		t.Errorf("Unexpected leaf. Expected %d, got %d", exp, act)
	}
}

func TestNew_WithFreshRoot_ShouldCreateValidTreeLeafRight(t *testing.T) {
	tr := New(5, 6)

	exp := 6
	act := tr.Root.Right.Value

	if act != exp {
		t.Errorf("Unexpected leaf. Expected %d, got %d", exp, act)
	}
}

func TestNew_WithFreshRoot_RootHasCorrectLength(t *testing.T) {
	tr := New(5, 6)

	exp := 1
	act := tr.Root.Length

	if act != exp {
		t.Errorf("Unexpected length. Expected %d, got %d", exp, act)
	}
}

func TestAdd_WithSingleNode_ShouldShiftRootNode(t *testing.T) {
	tr := New(5, 6)

	tr.Add(7, 8)

	expL := 5
	actL := tr.Root.Left.Left.Value

	if actL != expL {
		t.Errorf("Unexpected first leaf. Expected %d, got %d", expL, actL)
	}

	expR := 8
	actR := tr.Root.Right.Right.Value

	if actR != expR {
		t.Errorf("Unexpected first leaf. Expected %d, got %d", expL, actR)
	}
}

func TestAdd_WithSingleNode_ShouldHaveCorrectRootLength(t *testing.T) {
	tr := New(5, 6)

	tr.Add(7, 8)

	exp := 2
	act := tr.Root.Length

	if act != exp {
		t.Errorf("Unexpected length. Expected %d, got %d", exp, act)
	}
}

// func TestExplode_WithExplodableNode_ShouldAddValueToRegularLeftNode(t *testing.T) {
// 	tr := New(5, 6)

// }

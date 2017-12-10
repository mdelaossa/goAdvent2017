package puzzle1

import "testing"

func TestSolve(t *testing.T) {
	in := [][]int{
		[]int{5, 1, 9, 5},
		[]int{7, 5, 3},
		[]int{2, 4, 6, 8},
	}

	expected := 18

	solver := Solver{in}
	checksum := solver.Solve()
	if checksum != expected {
		t.Errorf("Expected Solve() to return %v but got %v instead", expected, checksum)
	}
}

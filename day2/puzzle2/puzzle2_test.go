package puzzle2

import "testing"

func TestSolve(t *testing.T) {
	in := [][]int{
		[]int{5, 9, 2, 8},
		[]int{9, 4, 7, 3},
		[]int{3, 8, 6, 5},
	}

	expected := 9

	solver := Solver{in}
	checksum := solver.Solve()
	if checksum != expected {
		t.Errorf("Expected Solve() to return %v but got %v instead", expected, checksum)
	}
}

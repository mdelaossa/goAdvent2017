package puzzle2

import "testing"

func TestSolve(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{1, 2},
		{12, 23},
		{23, 25},
		{747, 806},
	}

	for _, c := range cases {
		solver := Solver{}
		got := solver.Solve(c.in)
		if got != c.want {
			t.Errorf("Solve(%v) == %v, wanted: %v\n", c.in, got, c.want)
		}
	}
}

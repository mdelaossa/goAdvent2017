package puzzle1

import "testing"

func TestSolve(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{1, 0},
		{12, 3},
		{23, 2},
		{1024, 31},
	}

	for _, c := range cases {
		solver := Solver{}
		got := solver.Solve(c.in)
		if got != c.want {
			t.Errorf("Solve(%v) == %v, wanted: %v\n", c.in, got, c.want)
		}
	}
}

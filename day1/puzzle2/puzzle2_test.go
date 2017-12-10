package puzzle2

import "testing"

func TestEvaluatePair(t *testing.T) {
  cases := []struct {
    in, want []int
  } {
    {[]int{1, 2, 1, 2}, []int{1, 2, 1, 2}},
    {[]int{1, 2, 2, 1}, []int{}},
    {[]int{1, 2, 3, 4, 2, 5}, []int{2, 2}},
    {[]int{1, 2, 3, 1, 2, 3}, []int{1, 2, 3, 1, 2, 3}},
    {[]int{1, 2, 1, 3, 1, 4, 1, 5}, []int{1, 1, 1, 1}},
  }

  for _, c := range cases {
    got := evaluatePair(c.in)
    if len(got) != len(c.want) {
      t.Errorf("evaluatePair(%v) == %v, want %v", c.in, got, c.want)
    }
    for i, v := range got {
      if v != c.want[i] {
        t.Errorf("findAdjacent(%v) == %v, want %v", c.in, got, c.want)
      }
    }
  }
}

func TestSolve(t *testing.T) {
  cases := []struct {
    in []int
    want int
  } {
    {[]int{1, 2, 1, 2}, 6},
    {[]int{1, 2, 2, 1}, 0},
    {[]int{1, 2, 3, 4, 2, 5}, 4},
    {[]int{1, 2, 3, 1, 2, 3}, 12},
    {[]int{1, 2, 1, 3, 1, 4, 1, 5}, 4},
  }

  solver := Puzzle2Solver{}

  for _, c := range cases {
    got := solver.Solve(c.in)
    if got != c.want {
			t.Errorf("Solve(%v) == %v, want %v", c.in, got, c.want)
		}
  }
}

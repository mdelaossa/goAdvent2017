package puzzle1

import "testing"

func TestFindAdjacent(t *testing.T) {
  cases := []struct {
    in, want []int
  } {
    {[]int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
    {[]int{1, 1, 2, 2}, []int{1, 2}},
    {[]int{1, 2, 3, 4}, []int{}},
    {[]int{9, 1, 2, 1, 2, 1, 2, 9}, []int{9}},
  }

  for _, c := range cases {
    got := findAdjacent(c.in)
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
    {[]int{1, 1, 1, 1}, 4},
    {[]int{1, 1, 2, 2}, 3},
    {[]int{1, 2, 3, 4}, 0},
    {[]int{9, 1, 2, 1, 2, 1, 2, 9}, 9},
  }

  solver := Puzzle1Solver{}

  for _, c := range cases {
    got := solver.Solve(c.in)
    if got != c.want {
			t.Errorf("Solve(%v) == %v, want %v", c.in, got, c.want)
		}
  }
}

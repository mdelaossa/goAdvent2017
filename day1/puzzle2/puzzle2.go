package puzzle2

type Puzzle2Solver struct{}

func (s Puzzle2Solver) Solve(values []int) (sum int) {

  for _, v := range evaluatePair(values) {
    sum += v
  }

  return
}

func evaluatePair(values []int) (toSum []int) {
  toSum = make([]int, 0, len(values)/2)
  length := len(values)

  for i, val := range values {
    nextIndex := i + length/2
    if nextIndex >= length {
      nextIndex -= length
    }

    if val == values[nextIndex] {
      toSum = append(toSum, val)
    }
  }

  return
}

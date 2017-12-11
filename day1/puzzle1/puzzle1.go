package puzzle1

type Puzzle1Solver struct{}

func (s Puzzle1Solver) Solve(values []int) (sum int) {

	for _, v := range findAdjacent(values) {
		sum += v
	}

	return
}

func findAdjacent(values []int) (toSum []int) {
	toSum = make([]int, 0, len(values)/2)

	for i, val := range values {
		nextIndex := (i + 1) % len(values)

		if val == values[nextIndex] {
			toSum = append(toSum, val)
		}
	}

	return
}

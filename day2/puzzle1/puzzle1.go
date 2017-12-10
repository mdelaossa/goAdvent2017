package puzzle1

// Solver Solves Puzzle1 Day2
type Solver struct {
	Spreadsheet [][]int
}

// Solve computes the checksum for the given spreadsheet
func (p Solver) Solve() (checksum int) {
	lines := make([]int, len(p.Spreadsheet))

	for i, line := range p.Spreadsheet {
		min := int(^uint(0) >> 1) // set it to the max possible int value
		max := -min - 1           // set it to the minimum possible int value

		for _, cell := range line {
			if cell < min {
				min = cell
			}
			if cell > max {
				max = cell
			}
		}

		lines[i] = max - min
	}

	for _, v := range lines {
		checksum += v
	}

	return
}

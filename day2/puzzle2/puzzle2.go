package puzzle2

// Solver Solves Puzzle2 Day2
type Solver struct {
	Spreadsheet [][]int
}

// Solve computes the checksum for the given spreadsheet
func (p Solver) Solve() (checksum int) {
	lines := make([]int, len(p.Spreadsheet))

	for _, line := range p.Spreadsheet {
		// iterate from 0 to the penultimate cell
	Line:
		for i, a := range line[:len(line)-1] {
			// iterate from 1 to the last cell
			for b := i + 1; b < len(line); b++ {
				max, min := a, line[b]
				if max < min {
					max, min = min, max
				}
				if max%min == 0 {
					lines = append(lines, max/min)
					break Line
				}
			}
		}
	}

	for _, v := range lines {
		checksum += v
	}

	return
}

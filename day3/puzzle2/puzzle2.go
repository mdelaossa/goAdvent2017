package puzzle2

// Solver knows how to solze puzze 2 from Day 3
type Solver struct{}

// More convenient way to populate the spiral than using nested maps
type k struct {
	x, y int
}

// Solve actually solves puzze 2 from Day 3
func (s *Solver) Solve(input int) int {
	return buildSpiralUntil(input)
}

func buildSpiralUntil(n int) int {
	neighbors := [8]k{k{1, 0}, k{1, 1}, k{0, 1}, k{-1, 1}, k{-1, 0}, k{-1, -1}, k{0, -1}, k{1, -1}}
	directions := [4]k{k{1, 0}, k{0, 1}, k{-1, 0}, k{0, -1}}

	spiral := make(map[k]int)
	spiral[k{0, 0}] = 1
	x, y := 0, 0
	steps := 1
	secondDirection := false
	nstep := 0
	direction := 0

	for {
		v := directions[direction]
		x, y = x+v.x, y+v.y
		total := 0
		for _, neighbor := range neighbors {
			coord := k{neighbor.x + x, neighbor.y + y}
			if val, ok := spiral[coord]; ok {
				total += val
			}
		}

		if total > n {
			return total
		}
		spiral[k{x, y}] = total
		nstep++
		if nstep == steps {
			nstep = 0
			direction = (direction + 1) % 4
			if secondDirection {
				secondDirection = false
				steps++
			} else {
				secondDirection = true
			}
		}
	}
}

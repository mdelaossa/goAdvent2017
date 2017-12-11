package puzzle1

import (
	"math"
	"sync"
)

// Solver knows how to solze puzze 1 from Day 3
type Solver struct{}

// Solve actually solves puzze 1 from Day 3
func (s *Solver) Solve(input int) int {
	length := sideLength(input)
	stepsToLayer := (length - 1) / 2
	midpoint := closestMidpoint(input, length)
	stepsToMidpoint := input - midpoint
	if stepsToMidpoint < 0 {
		stepsToMidpoint = -stepsToMidpoint
	}
	return stepsToLayer + stepsToMidpoint
}

func sideLength(input int) int {
	fl := float64(input)
	tmp := math.Ceil(math.Sqrt(fl))
	res := int(tmp)
	if res%2 == 0 {
		return res + 1
	}
	return res
}

func closestMidpoint(input int, length int) int {
	highestValue := length * length
	offset := int((length - 1) / 2.0)
	closestMidpoint := 0

	wg := sync.WaitGroup{}
	ch := make(chan int)

	for i := 0; i < 4; i++ {
		wg.Add(1)
		// When you have a shiny new hammer, everything looks like a nail
		go func(wg *sync.WaitGroup, ch chan int, iteration int) {
			defer wg.Done()
			midpoint := highestValue - (offset * (2*iteration + 1))
			ch <- midpoint
		}(&wg, ch, i)
	}

	go func(wg *sync.WaitGroup, ch chan int) {
		wg.Wait()
		close(ch)
	}(&wg, ch)

	for midpoint := range ch {
		dst := input - midpoint
		if dst < 0 {
			dst = -dst
		}

		closestDst := input - closestMidpoint
		if closestDst < 0 {
			closestDst = -closestDst
		}

		if dst < closestDst {
			closestMidpoint = midpoint
		}
	}

	return closestMidpoint
}

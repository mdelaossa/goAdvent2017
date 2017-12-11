package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/mdelaossa/goAdvent2017/day3/puzzle1"
	"github.com/mdelaossa/goAdvent2017/day3/puzzle2"
)

type spiralSolver interface {
	Solve(int) int
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("ERROR: Missing argumnts. Proper usage: ./day3 <puzzle> <input>")
	}

	solver := chooseSolver(os.Args[1])

	input, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("ERROR: Input must be a number! (Got: %v)", os.Args[2])
	}

	fmt.Printf("Your solution is: %v", solver.Solve(input))
}

func chooseSolver(puzzle string) (solver spiralSolver) {
	switch puzzle {
	case "1":
		solver = &puzzle1.Solver{}
	case "2":
		solver = &puzzle2.Solver{}
	}
	return
}

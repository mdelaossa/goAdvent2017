package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mdelaossa/goAdvent2017/day2/puzzle1"
	"github.com/mdelaossa/goAdvent2017/day2/puzzle2"
)

type checksumSolver interface {
	Solve() int
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("ERROR: Missing argument: puzzle (1 or 2)")
	}

	fmt.Println("Please input the spreadsheet separating each line with a newline.")
	fmt.Println("Send EOF (ctrl+d) on it's own line when done.")
	values, err := scanInts()
	if err != nil {
		log.Fatal("Error: Could not get input:", err)
	}

	solver := selectSolver(os.Args[1], values)

	fmt.Printf("The checksum is: %v\n", solver.Solve())
}

func scanInts() ([][]int, error) {
	scanner := bufio.NewScanner(os.Stdin)

	ints := make([][]int, 0)

	for scanner.Scan() {
		current := make([]int, 0)

		for _, str := range strings.Fields(scanner.Text()) {
			i, err := strconv.Atoi(str)
			if err != nil {
				return ints, err
			}
			current = append(current, i)
		}
		ints = append(ints, current)
	}
	return ints, scanner.Err()
}

func selectSolver(puzzle string, values [][]int) (solver checksumSolver) {
	switch puzzle {
	case "1":
		solver = puzzle1.Solver{values}
	case "2":
		solver = puzzle2.Solver{values}
	}
	return
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/mdelaossa/goAdvent2017/day4/puzzle1"
	"github.com/mdelaossa/goAdvent2017/day4/puzzle2"
)

type passSolver interface {
	Solve(*[]string) int
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("ERROR: Missing arguments. Proper usage: ./day4 <puzzle>")
	}

	solver := chooseSolver(os.Args[1])

	fmt.Println("Please input the passphrases separating each with a newline.")
	fmt.Println("Send EOF (ctrl+d) on it's own line when done.")
	input, err := scanPhrases()
	if err != nil {
		log.Fatalf("ERROR: could not receive input: %v", err)
	}

	fmt.Printf("Your solution is: %v", solver.Solve(&input))
}

func scanPhrases() ([]string, error) {
	scanner := bufio.NewScanner(os.Stdin)

	phrases := make([]string, 0, 10)

	for scanner.Scan() {
		phrases = append(phrases, scanner.Text())
	}

	return phrases, scanner.Err()
}

func chooseSolver(puzzle string) (solver passSolver) {
	switch puzzle {
	case "1":
		solver = &puzzle1.Solver{}
	case "2":
		solver = &puzzle2.Solver{}
	}
	return
}

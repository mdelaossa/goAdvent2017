package main

import (
  "fmt"
  "strconv"
  "strings"
  "os"
  "log"
  "github.com/mdelaossa/goAdvent2017/day1/puzzle1"
  "github.com/mdelaossa/goAdvent2017/day1/puzzle2"
)

type captchaSolver interface {
  Solve([]int) int
}

func main() {

  if len(os.Args) < 2 {
    log.Fatal("ERROR: Missing argument: puzzle (1 or 2)")
  }

  fi, _ := os.Stdin.Stat()
  if (fi.Mode() & os.ModeCharDevice) != 0 {
    log.Fatal("ERROR: Due to golang constraints the input to the captcha solver must be piped in as it's too large. Ex. `echo 1111 | ./day1 1`")
  }

  var input string
  fmt.Scan(&input)

  numbers := prepareData(input)

  solver := selectPuzzle(os.Args[1])

  fmt.Printf("The solution to the captcha is: %v\n", solver.Solve(numbers))
}

func selectPuzzle(input string) (solver captchaSolver) {
  switch input {
  case "1":
    solver = puzzle1.Puzzle1Solver{}
  case "2":
    solver = puzzle2.Puzzle2Solver{}
  }

  return
}

func prepareData(input string) (numbers []int) {
  split := strings.Split(input, "")
  numbers = make([]int, 0, len(split))

  for _, str := range split {
    val, err := strconv.Atoi(str)
    if err != nil {
      panic(err)
    }
    numbers = append(numbers, val)
  }

  return
}

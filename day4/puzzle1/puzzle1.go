package puzzle1

import (
	"sort"
	"strings"
)

// Solver struct
type Solver struct{}

// Solve actually solves the puzzle
func (s Solver) Solve(phrases *[]string) (validCount int) {
	for _, phrase := range *phrases {
		if isValid(phrase) {
			validCount++
		}
	}

	return
}

func isValid(phrase string) bool {
	strs := strings.Split(phrase, " ")
	sort.Strings(strs)
	for i, str := range strs {
		if i+1 == len(strs) {
			break
		}
		if str == strs[i+1] {
			return false
		}
	}
	return true
}

package puzzle2

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
	for i, str := range strs[:len(strs)-1] {
		for _, str2 := range strs[i+1:] {
			if isAnagram(str, str2) {
				return false
			}
		}
	}
	return true
}

func isAnagram(word1 string, word2 string) bool {
	word1 = sortString(word1)
	word2 = sortString(word2)

	return word1 == word2
}

func sortString(s string) string {
	sl := strings.Split(s, "")
	sort.Strings(sl)
	return strings.Join(sl, "")
}

package puzzle2

import "testing"

func TestIsValid(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"abcde fghij", true},
		{"abcde xyz ecdab", false},
		{"a ab abc abd abf abj", true},
		{"iiii oiii ooii oooi oooo", true},
		{"oiii ioii iioi iiio", false},
	}

	for _, c := range cases {
		got := isValid(c.in)
		if got != c.want {
			t.Errorf("isValid(%v) == %v, wanted: %v\n", c.in, got, c.want)
		}
	}
}

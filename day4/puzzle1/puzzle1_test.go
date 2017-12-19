package puzzle1

import "testing"

func TestIsValid(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"aa bb cc dd ee", true},
		{"aa bb cc dd aa", false},
		{"aa bb cc dd aaa", true},
	}

	for _, c := range cases {
		got := isValid(c.in)
		if got != c.want {
			t.Errorf("isValid(%v) == %v, wanted: %v\n", c.in, got, c.want)
		}
	}
}

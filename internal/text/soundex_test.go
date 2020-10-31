package text

import (
	"testing"
)

func TestSoundex(t *testing.T) {
	type testCase struct {
		st       string
		expected string
	}
	cases := []testCase{
		{st: "I'm", expected: "I500"},
		{st: "I am", expected: "I500"},
		{st: "am", expected: "A500"}}
	for _, c := range cases {
		out := Soundex(c.st)
		if out != c.expected {
			t.Errorf("Failed: expected %v, recieved %v", c.expected, out)
		}
	}
}

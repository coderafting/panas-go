package text

import (
	"testing"
)

func TestGenerateValidWords(t *testing.T) {
	testCase := "I'm happy\n now! "
	expected := []string{"I'm", "happy", "now!"}
	out := GenerateValidWords(testCase)
	if out[0] != expected[0] && out[1] != expected[1] && out[2] != expected[2] {
		t.Errorf("Failed: expected %v, recieved %v", expected, out)
	}
}

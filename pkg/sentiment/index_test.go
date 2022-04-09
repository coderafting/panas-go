package sentiment

import (
	"testing"
)

func TestBuildSoundexIndex(t *testing.T) {
	testCase := []string{"I'm", "I am", "am"}
	expected := map[string][]string{"I500": {"I'm", "I am"}, "A500": {"am"}}
	out := BuildSoundexIndex(testCase)
	if out["I500"][0] != expected["I500"][0] && out["I500"][1] != expected["I500"][1] && out["A500"][0] != expected["A500"][0] {
		t.Errorf("Failed: expected %v, recieved %v", expected, out)
	}
}

package sentiment

import (
	"testing"
)

func TestContainsTopic(t *testing.T) {
	type testCase struct {
		topic    string
		words    []string
		expected bool
	}
	cases := []testCase{
		{topic: "Covid", words: []string{"this", "is", "covid", "time"}, expected: true},
		{topic: "Corona", words: []string{"this", "is", "covid", "time"}, expected: false}}
	for _, c := range cases {
		out := ContainsTopic(c.topic, c.words)
		if out != c.expected {
			t.Errorf("Failed: expected %v, recieved %v", c.expected, out)
		}
	}
}

func TestInIndex(t *testing.T) {
	type testCase struct {
		indexMap map[string][]string
		words    []string
		expected bool
	}
	index := map[string][]string{"I500": {"I'm", "I am"}, "A500": {"am"}}
	cases := []testCase{
		{indexMap: index, words: []string{"I am", "good"}, expected: true},
		{indexMap: index, words: []string{"this", "is", "covid", "time"}, expected: false}}

	for _, c := range cases {
		out := InIndex(c.indexMap, c.words)
		if out != c.expected {
			t.Errorf("Failed: expected %v, recieved %v", c.expected, out)
		}
	}
}

func TestContainsOneSelfRef(t *testing.T) {
	type testCase struct {
		words    []string
		expected bool
	}
	cases := []testCase{
		{words: []string{"I am", "good"}, expected: true},
		{words: []string{"this", "is", "covid", "time"}, expected: false}}

	for _, c := range cases {
		out := ContainsOneSelfRef(c.words)
		if out != c.expected {
			t.Errorf("Failed: expected %v, recieved %v", c.expected, out)
		}
	}
}

func TestContainsValidSentiment(t *testing.T) {
	type testCase struct {
		words    []string
		expected bool
	}
	cases := []testCase{
		{words: []string{"I am", "happy"}, expected: true},
		{words: []string{"this", "is", "covid", "time"}, expected: false}}

	for _, c := range cases {
		out := ContainsValidSentiment(c.words)
		if out != c.expected {
			t.Errorf("Failed: expected %v, recieved %v", c.expected, out)
		}
	}
}

func TestValidText(t *testing.T) {
	type testCase struct {
		textString string
		expected   bool
	}
	cases := []testCase{
		{textString: "I am happy", expected: true},
		{textString: "I am wondering", expected: false},
		{textString: "this is covid time", expected: false},
		{textString: "this is sad time", expected: false}}

	for _, c := range cases {
		out := ValidText(c.textString)
		if out != c.expected {
			t.Errorf("Failed: expected %v, recieved %v", c.expected, out)
		}
	}
}

func TestValidTextWithTopic(t *testing.T) {
	type testCase struct {
		textString string
		topic      string
		expected   bool
	}
	cases := []testCase{
		{textString: "I am happy that covid is getting under control", topic: "covid", expected: true},
		{textString: "I am happy", topic: "covid", expected: false},
		{textString: "I am wondering", topic: "covid", expected: false},
		{textString: "this is covid time", topic: "covid", expected: false},
		{textString: "this is sad time", topic: "covid", expected: false}}

	for _, c := range cases {
		out := ValidTextWithTopic(c.textString, c.topic)
		if out != c.expected {
			t.Errorf("Failed: expected %v, recieved %v", c.expected, out)
		}
	}
}

func TestStates(t *testing.T) {
	type testCase struct {
		textString string
		expected   []string
	}
	cases := []testCase{
		{textString: "I am xyz", expected: []string{}},
		{textString: "I am happy", expected: []string{"happy"}},
		{textString: "I am both happy and sad", expected: []string{"happy", "sad"}}}

	for _, c := range cases {
		out := States(c.textString)
		if len(out) == 1 {
			if out[0] != c.expected[0] {
				t.Errorf("Failed: expected %v, recieved %v", c.expected, out)
			}
		} else if len(out) > 1 {
			if out[0] != c.expected[0] || out[1] != c.expected[1] {
				t.Errorf("Failed: expected %v, recieved %v", c.expected, out)
			}
		} else if len(out) == 0 {
			if len(c.expected) != 0 {
				t.Errorf("Failed: expected %v, recieved %v", c.expected, out)
			}
		}
	}
}

func TestCategories(t *testing.T) {
	type testCase struct {
		textString string
		expected   []string
	}
	cases := []testCase{
		{textString: "I am xyz", expected: []string{}},
		{textString: "I am happy", expected: []string{"jovility"}},
		{textString: "I am happy, joyful, and sad", expected: []string{"jovility", "sadness"}}}

	for _, c := range cases {
		out := Categories(c.textString)
		if len(out) == 1 {
			if out[0] != c.expected[0] {
				t.Errorf("Failed: expected %v, recieved %v", c.expected, out)
			}
		} else if len(out) > 1 {
			if out[0] != c.expected[0] || out[1] != c.expected[1] {
				t.Errorf("Failed: expected %v, recieved %v", c.expected, out)
			}
		} else if len(out) == 0 {
			if len(c.expected) != 0 {
				t.Errorf("Failed: expected %v, recieved %v", c.expected, out)
			}
		}
	}
}

func TestCategoryAggregate(t *testing.T) {
	type testCase struct {
		categoryTextsCount int
		totalTextsCount    int
		expected           float64
	}
	cases := []testCase{
		{categoryTextsCount: 5, totalTextsCount: 10, expected: 0.5},
		{categoryTextsCount: 2, totalTextsCount: 10, expected: 0.2}}

	for _, c := range cases {
		out, _ := CategoryAggregate(c.categoryTextsCount, c.totalTextsCount)
		if out != c.expected {
			t.Errorf("Failed: expected %v, recieved %v", c.expected, out)
		}
	}
}

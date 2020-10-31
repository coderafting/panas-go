// Package sentiment exposes APIs to extract sentiment category of a text and compute aggregate sentiment values.
// The implementation is based on the paper "PANAS-t: A Pychometric Scale for Measuring Sentiments on Twitter",
// which can be accessed at https://arxiv.org/abs/1308.1857.
package sentiment

import (
	"github.com/coderafting/panas-go/internal/text"
)

/*
Text validation, Extract sentiment types and Compute aggregate sentiment value for texts.
Based on the PANAS-t paper.
*/

// ContainsTopic checks if the words-collection contains at least one word
// that is similar to the supplied word (topic).
func ContainsTopic(topic string, words []string) bool {
	topicSoundex := text.Soundex(topic)
	for _, w := range words {
		if topicSoundex == text.Soundex(w) {
			return true
		}
	}
	return false
}

// InIndex checks if the words-collection contains at least one word that exists in the supplied index map.
func InIndex(indexMap map[string][]string, words []string) bool {
	for _, w := range words {
		if indexMap[text.Soundex(w)] != nil {
			return true
		}
	}
	return false
}

// ContainsOneSelfRef checks if the words-collection contains at least one word that is similar to
// one of the selfReferences recognized by the PANAS-t paper.
func ContainsOneSelfRef(words []string) bool {
	return InIndex(SelfRefSoundexIndex, words)
}

// ContainsValidSentiment checks if the words-collection contains at least one word that is similar to
// one of the sentimentStates recognized by the PANAS-t paper.
func ContainsValidSentiment(words []string) bool {
	return InIndex(StatesSoundexIndex, words)
}

// ValidText returns true if it finds the text to be valid to be considered for sentiment analysis.
// Means, the text must contain a self reference and a sentiment state.
func ValidText(textString string) bool {
	words := text.GenerateValidWords(textString)
	if ContainsOneSelfRef(words) && ContainsValidSentiment(words) {
		return true
	}
	return false
}

// ValidTextWithTopic returns true if it finds the text to be valid to be considered for sentiment analysis on a topic.
// Means, the text must contain the target topic, a self reference, and a sentiment state.
func ValidTextWithTopic(textString, topic string) bool {
	words := text.GenerateValidWords(textString)
	if ContainsTopic(topic, words) && ContainsOneSelfRef(words) && ContainsValidSentiment(words) {
		return true
	}
	return false
}

// Categories detrmines the sentiment categories of a text.
func Categories(textString string) []string {
	// Currently, the conflict resolution is ignored when a tweet contains more than one sentiment.
	// For now, we will simply consider such tweets a part of all identified sentiment categories.
	catgs := []string{}
	words := text.GenerateValidWords(textString)
	for _, w := range words {
		states := StatesSoundexIndex[text.Soundex(w)]
		if states != nil {
			for _, s := range states {
				// NOTE: the states contains 1 or 2 items, therefore, this nested for-loop is perfectly alright.
				stateCatg := StatesCategories[s].category
				catgs = append(catgs, stateCatg)
			}
		}
	}
	return catgs
}

// CategoryAggregate returns the aggregate sentiment value of a category. It ranges from 0 to 1.
func CategoryAggregate(categoryTextsCount int, totalTextsCount int) float64 {
	return float64(categoryTextsCount) / float64(totalTextsCount)
}

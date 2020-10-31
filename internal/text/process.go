// Package text exposes utility functions to process a text string.
package text

import (
	"regexp"
	"strings"
)

/*
Create text vecs, and check if it is a valid text based on subject, topic, and sentiment text
*/

// GenerateValidWords returns a list of words processed from the input text.
func GenerateValidWords(text string) []string {
	wordsColl := strings.Split(text, " ")
	validWords := []string{}
	for _, w := range wordsColl {
		lowerCased := strings.ToLower(w)
		newLineRemoved := strings.ReplaceAll(lowerCased, "\n", "")
		spacesTrimmed := strings.TrimSpace(newLineRemoved)
		reg := regexp.MustCompile("[^A-Za-z0-9]+")
		processedWord := reg.ReplaceAllString(spacesTrimmed, "")
		validWords = append(validWords, processedWord)
	}
	return validWords
}

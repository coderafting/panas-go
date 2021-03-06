package text

/*
Implementation of Soundex algorithm: https://en.wikipedia.org/wiki/Soundex
This implementation code is copied directly from https://github.com/xrash/smetrics/blob/master/soundex.go
*/

import (
	"strings"
)

// Soundex encoding is a phonetic algorithm that considers how the words sound in english.
// Soundex maps a name to a 4-byte string consisting of the first letter of the original string and three numbers.
// Strings that sound similar should map to the same thing.
func Soundex(s string) string {
	m := map[byte]string{
		'B': "1", 'P': "1", 'F': "1", 'V': "1",
		'C': "2", 'S': "2", 'K': "2", 'G': "2", 'J': "2", 'Q': "2", 'X': "2", 'Z': "2",
		'D': "3", 'T': "3",
		'L': "4",
		'M': "5", 'N': "5",
		'R': "6",
	}
	s = strings.ToUpper(s)
	r := string(s[0])
	p := s[0]
	for i := 1; i < len(s) && len(r) < 4; i++ {
		c := s[i]
		if (c < 'A' || c > 'Z') || (c == p) {
			continue
		}
		p = c
		if n, ok := m[c]; ok {
			r += n
		}
	}
	for i := len(r); i < 4; i++ {
		r += "0"
	}
	return r
}

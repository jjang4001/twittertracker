package parser

import (
	"fmt"
	"strings"
	"unicode"
)

// GetWords takes a tweet string and removes stopwords
func GetWords(tweetText string) []string {
	var words = strings.Split(tweetText, " ")
	var finalWords []string
	for i := 0; i < len(words); i++ {
		word := words[i]

		// Pre-trim processing
		if strings.ContainsRune(word, '@') ||
			strings.Contains(word, "http") ||
			strings.Contains(word, "…") ||
			strings.Contains(word, "...") {
			continue
		}

		// Trim
		word = strings.TrimFunc(strings.ToLower(words[i]), isNonLetter)

		// Post-trim processing
		if !Stopwords[word] {
			fmt.Println(word)
			finalWords = append(finalWords, word)
		}
	}
	return finalWords
}

func isNonLetter(r rune) bool {
	return !unicode.IsLetter(r)
}

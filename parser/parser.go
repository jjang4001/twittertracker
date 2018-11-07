package parser

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/dghubble/go-twitter/twitter"
)

// GetWordsFromTweet takes a tweet and tries to parse it as a tweet, removes stopwords, and returns a list of words
func GetWordsFromTweet(tweet twitter.Tweet) []string {
	// perform task
	fmt.Println("Processing a tweet")
	if tweet.RetweetedStatus != nil {
		fmt.Println("Retweeted")
		return retrieveCorrectType(*tweet.RetweetedStatus)
	}
	return retrieveCorrectType(tweet)
}

func getWords(tweetText string) []string {
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

func retrieveCorrectType(tweet twitter.Tweet) []string {
	if tweet.ExtendedTweet != nil {
		fmt.Println("Extended")
		return PrintAndParse(tweet.ExtendedTweet.FullText)
	}
	fmt.Println("Normal")
	return PrintAndParse(tweet.Text)
}

// PrintAndParse retrieves the normalized word bag
func PrintAndParse(text string) []string {
	fmt.Println(text)
	fmt.Println("-------------------------------------")
	words := getWords(strings.Replace(text, "\n", " ", -1))
	fmt.Println(words)
	fmt.Println("=======================================")
	return words
}

func isNonLetter(r rune) bool {
	return !unicode.IsLetter(r)
}

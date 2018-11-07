package parser

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/dghubble/go-twitter/twitter"
)

// GetWordsFromTweet takes a tweet and tries to parse it as a tweet, removes stopwords, and returns a list of words
func GetWordsFromTweet(tweet twitter.Tweet, parsedWords chan<- string) {
	// perform task
	fmt.Println("Processing a tweet")
	if tweet.RetweetedStatus != nil {
		fmt.Println("Retweeted")
		retrieveCorrectType(*tweet.RetweetedStatus, parsedWords)
		return
	}
	retrieveCorrectType(tweet, parsedWords)
}

func getWords(tweetText string, parsedWords chan<- string) {
	var words = strings.Split(tweetText, " ")

	defer close(parsedWords)
	for _, word := range words {

		// Pre-trim processing
		if strings.ContainsRune(word, '@') ||
			strings.Contains(word, "http") ||
			strings.Contains(word, "…") ||
			strings.Contains(word, "...") {
			continue
		}

		// Trim
		word = strings.TrimFunc(strings.ToLower(word), isNonLetter)

		// Post-trim processing
		if !Stopwords[word] {
			fmt.Println("sending word", word)
			parsedWords <- word
		}
	}
}

func retrieveCorrectType(tweet twitter.Tweet, parsedWords chan<- string) {
	if tweet.ExtendedTweet != nil {
		fmt.Println("Extended")
		PrintAndParse(tweet.ExtendedTweet.FullText, parsedWords)
		return
	}
	fmt.Println("Normal")
	PrintAndParse(tweet.Text, parsedWords)
}

// PrintAndParse retrieves the normalized word bag
func PrintAndParse(text string, parsedWords chan<- string) {
	fmt.Println("PrintAndParse: ", text)
	fmt.Println("-------------------------------------")
	getWords(strings.Replace(text, "\n", " ", -1), parsedWords)
	fmt.Println("=======================================")
}

func isNonLetter(r rune) bool {
	return !unicode.IsLetter(r)
}

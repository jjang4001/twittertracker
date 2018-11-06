package parser

import (
	"encoding/json"
	"fmt"
	"strings"
	"unicode"

	"github.com/dghubble/go-twitter/twitter"
)

// GetWordsFromTweet takes a string payload tries to parse it as a tweet, removes stopwords, and returns a list of words
func GetWordsFromTweet(payload string, words []string) error {
	var tweet twitter.Tweet
	var finalStrings []string
	fmt.Println(payload)
	if err := json.Unmarshal([]byte(payload), &tweet); err != nil {
		// handle error
		return err
	}

	// perform task
	fmt.Println("Processing a tweet")
	fmt.Println(tweet.ExtendedTweet)
	if tweet.RetweetedStatus != nil {
		fmt.Println("Retweeted")
		finalStrings = retrieveCorrectType(*tweet.RetweetedStatus)
	} else {
		finalStrings = retrieveCorrectType(tweet)
	}
	words = finalStrings[:]
	fmt.Println("=======================================")
	return nil
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
	return words
}

func isNonLetter(r rune) bool {
	return !unicode.IsLetter(r)
}

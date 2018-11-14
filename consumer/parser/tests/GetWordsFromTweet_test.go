package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
	"twittertracker/parser"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/stretchr/testify/assert"
)

func Test_GetWordsFromTweet(t *testing.T) {
	assert := assert.New(t)

	dat, err := ioutil.ReadFile("tweet.json")
	assert.Equal(nil, err)

	var tweet twitter.Tweet
	if err := json.Unmarshal([]byte(dat), &tweet); err != nil {
		assert.FailNow(err.Error())
	}

	parsedWords := make(chan string)
	expected := []string{"president", "landed", "hoosier", "state", "tax", "cuts", "soaring", "median", "income", "indianans", "thriving", "trump", "economy"}
	go parser.GetWordsFromTweet(tweet, parsedWords)
	receiveWordsFromChannel(parsedWords, expected, t)
}

func receiveWordsFromChannel(parsedWords <-chan string, expected []string, t *testing.T) {
	fmt.Println("running receiveWordsFromChannel")
	assert := assert.New(t)
	numReceived := 0
	for word := range parsedWords {
		assert.Equal(word, expected[numReceived])
		numReceived++
		fmt.Println("CHANNEL: ", word)
	}
}

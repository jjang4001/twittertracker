package main

import (
	"encoding/json"
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

	words := parser.GetWordsFromTweet(tweet)
	expected := []string{"president", "landed", "hoosier", "state", "tax", "cuts", "soaring", "median", "income", "indianans", "thriving", "trump", "economy"}
	assert.Equal(expected, words)
}

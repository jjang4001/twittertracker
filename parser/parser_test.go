package parser

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_GetWordsFromTweet is a unit tests that ensures the json gets the correct words
func Test_GetWordsFromTweet(t *testing.T) {
	var words = make([]string, 1)
	assert := assert.New(t)

	dat, err := ioutil.ReadFile("test.json")
	assert.Equal(nil, err)

	GetWordsFromTweet(string(dat), words)
	expected := []string{"president", "landed", "hoosier", "state", "tax", "cuts", "soaring", "median", "income", "indianans", "thriving", "trump", "economy"}
	assert.Equal(expected, words)
}

func Test_ParseString(t *testing.T) {
	// assert := assert.New(t)
	// input := "A little more than 24 hours from now Trump will declare victory.\n\nIf Republicans do well, Trump will take credit\n\nIf Democrats do well Trump will say it would have been much worse without his rallies. He will also blame “illegal” voters and the media.\n\nBookmark this."
	// PrintAndParse(input)
	// expected := []string{"little", "hours", "trump", "declare", "victory", "republicans", "well", "trump", "take", "credit", "democrats", "well", "trump", "say", "would", "much", "worse", "without", "rallies", "also", "blame", "illegal", "voters", "media", "bookmark"}
	// assert.Equal(expected, output)
}

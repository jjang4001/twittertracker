package tests

import (
	"testing"
	"twittertracker/parser"
)

func Test_ParseString(t *testing.T) {
	// assert := assert.New(t)
	dummyChan := make(chan string)
	input := "A little more than 24 hours from now Trump will declare victory.\n\nIf Republicans do well, Trump will take credit\n\nIf Democrats do well Trump will say it would have been much worse without his rallies. He will also blame “illegal” voters and the media.\n\nBookmark this."
	parser.PrintAndParse(input, dummyChan)
	// expected := []string{"little", "hours", "trump", "declare", "victory", "republicans", "well", "trump", "take", "credit", "democrats", "well", "trump", "say", "would", "much", "worse", "without", "rallies", "also", "blame", "illegal", "voters", "media", "bookmark"}
}

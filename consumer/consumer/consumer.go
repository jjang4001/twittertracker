package consumer

import (
	"encoding/json"
	"fmt"
	"twittertracker/datastore"
	"twittertracker/parser"

	"github.com/adjust/rmq"
	"github.com/dghubble/go-twitter/twitter"
)

// TaskConsumer implements the Consumer interface
type TaskConsumer struct {
	DbEnv *datastore.Env
}

// Consume is the work that will be done by the taskconsumer
func (consumer *TaskConsumer) Consume(delivery rmq.Delivery) {
	var tweet twitter.Tweet

	fmt.Println("Printing JSON")
	payload := delivery.Payload()
	fmt.Println(payload)

	if err := json.Unmarshal([]byte(payload), &tweet); err != nil {
		// handle error
		delivery.Reject()
		return
	}

	parsedWords := make(chan string)
	go parser.GetWordsFromTweet(tweet, parsedWords)
	saveWord(parsedWords, consumer.DbEnv)
	delivery.Ack()
}

func saveWord(parsedWords <-chan string, dbEnv *datastore.Env) {
	for word := range parsedWords {
		fmt.Println("saving word:", word)
		dbEnv.DB.SaveWord(word)
	}
}

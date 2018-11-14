package consumer

import (
	"encoding/json"
	"fmt"
	"twittertracker/consumer/parser"
	"twittertracker/datastore"

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
	if consumer.DbEnv.DB.BeginTransaction() != nil {
		delivery.Reject()
	}
	saveWord(parsedWords, consumer.DbEnv)
	if consumer.DbEnv.DB.ExecTransaction() != nil {
		delivery.Reject()
	}
	fmt.Println("Successfully saved to redis")
	delivery.Ack()
}

func saveWord(parsedWords <-chan string, dbEnv *datastore.Env) {
	for word := range parsedWords {
		fmt.Println("Adding to transaction:", word)
		dbEnv.DB.SaveWord(word)
	}
}

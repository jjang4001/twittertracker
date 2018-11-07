package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"twittertracker/common"
	"twittertracker/parser"

	"github.com/adjust/rmq"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/joho/godotenv"
)

// TaskConsumer implements the Consumer interface
type taskConsumer struct {
}

// Consume is the work that will be done by the taskconsumer
func (consumer *taskConsumer) Consume(delivery rmq.Delivery) {
	var tweet twitter.Tweet

	fmt.Println("Printing JSON")
	payload := delivery.Payload()
	fmt.Println(payload)

	if err := json.Unmarshal([]byte(payload), &tweet); err != nil {
		// handle error
		delivery.Reject()
		return
	}
	parser.GetWordsFromTweet(tweet)
	delivery.Ack()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Getting Queue")
	dbConnectionString := os.Getenv("LOCAL_REDIS")
	redisConn := rmq.OpenConnection(common.RedisQueueTag, common.RedisQueueProtocol, dbConnectionString, common.RedisQueueDB)
	taskQueue := redisConn.OpenQueue(common.RedisQueueName)
	taskQueue.StartConsuming(100, time.Second)

	taskConsumer := &taskConsumer{}
	taskQueue.AddConsumer("task consumer", taskConsumer)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)
	fmt.Println("Stopping Processing...")
}

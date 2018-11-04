package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adjust/rmq"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/joho/godotenv"
)

// TaskConsumer implements the Consumer interface
type TaskConsumer struct {
}

// Consume is the work that will be done by the taskconsumer
func (consumer *TaskConsumer) Consume(delivery rmq.Delivery) {
	var tweet twitter.Tweet
	if err := json.Unmarshal([]byte(delivery.Payload()), &tweet); err != nil {
		// handle error
		delivery.Reject()
		return
	}

	// perform task
	fmt.Println("performing task", tweet.User, tweet.IDStr)
	fmt.Println(tweet.Text)
	fmt.Println()
	delivery.Ack()
}

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Getting Queue")
	dbConnectionString := os.Getenv("LOCAL_REDIS")
	redisConn := rmq.OpenConnection("queue", "tcp", dbConnectionString, 1)
	taskQueue := redisConn.OpenQueue("tweets")
	taskQueue.StartConsuming(100, time.Second)

	taskConsumer := &TaskConsumer{}
	taskQueue.AddConsumer("task consumer", taskConsumer)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)
	fmt.Println("Stopping Processing...")
}

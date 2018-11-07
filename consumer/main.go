package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"twittertracker/common"
	"twittertracker/consumer/consumer"
	"twittertracker/datastore"

	"github.com/adjust/rmq"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Getting Queue")
	dbConnectionString := os.Getenv(common.LOCAL_REDIS)
	redisConn := rmq.OpenConnection(common.RedisQueueTag, common.RedisQueueProtocol, dbConnectionString, common.RedisQueueDB)
	taskQueue := redisConn.OpenQueue(common.RedisQueueName)
	taskQueue.StartConsuming(100, time.Second)

	db, err := datastore.NewDatastore(datastore.REDIS, dbConnectionString)
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	env := datastore.Env{DB: db}
	taskConsumer := &consumer.TaskConsumer{DbEnv: &env}
	taskQueue.AddConsumer(common.TASK_CONSUMER, taskConsumer)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)
	fmt.Println("Stopping Processing...")
}

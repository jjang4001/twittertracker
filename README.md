# Twitter Tracker

## Set up

1. Clone into ~/go/src directory
2. For each dependency, run:

```
go get <dependency>
ex: go get github.com/gorilla/handlers
```

1. To start the server, run


```
go run main.go
```

To test the websocket connection, open socket.html, and click "Test socket".

## Setting Up Redis

Steps to run:

Set up local redis server. Follow instructions at https://redis.io/topics/quickstart up to the section that says "Check if Redis is Working"

go get all the new dependencies.

Create .env file in root directory with
```
LOCAL_REDIS=<Your local redis host server> (i.e. 127.0.0.1:6379)
PORT=:3000 (or any port you want)
```
Make sure your redis server is running.
run "go run main.go", and then send a POST request to "http://localhost:3000/example/mykey"
send GET request to the same endpoint, and hopefully everything works.

## Setting Up Twitter

Add the following keys in the .env file

```
CONSUMER_KEY=<Account Consumer Key>
CONSUMER_SECRET=<Account Consumer Secret>
ACCESS_KEY=<Account Access Key>
ACCESS_SECRET=<Account Access Secret>
```

### Producer
To start the producer go to the producer folder and run `go run producer.go`. You should see the following buffer start to flood the screen: `Published tweet <TweetId>`

### Consumer
In a separate terminal window navigate to the consumer folder and run `go run consumer.go` to see the consumer reading events from the redis queue.
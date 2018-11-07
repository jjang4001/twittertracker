package common

// RedisQueueDB is the associated redis database number
const RedisQueueDB = 1

// RedisQueueTag is the tag used to
const RedisQueueTag = "queue"

// RedisQueueName is the name of the queue of tweets
const RedisQueueName = "tweets"

// RedisQueueProtocol is the protocol used for connecting to the redis queue
const RedisQueueProtocol = "tcp"

// Redis message queue constants
const TaskConsumer = "tweet consumer"

// Local environment variables
const LocalRedis = "LOCAL_REDIS"
const Port = "PORT"
const ConsumerKey = "CONSUMER_KEY"
const ConsumerSecret = "CONSUMER_SECRET"
const AccessKey = "ACCESS_KEY"
const AccessSecret = "ACCESS_SECRET"

// Redis commands
const Set = "SET"
const Exists = "EXISTS"
const Get = "GET"
const Incr = "INCR"
const Multi = "MULTI"
const Exec = "EXEC"

// Handler methods
const Post = "POST"

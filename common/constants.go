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
const TASK_CONSUMER = "tweet consumer"

// Local environment variables
const LOCAL_REDIS = "LOCAL_REDIS"
const PORT = "PORT"
const CONSUMER_KEY = "CONSUMER_KEY"
const CONSUMER_SECRET = "CONSUMER_SECRET"
const ACCESS_KEY = "ACCESS_KEY"
const ACCESS_SECRET = "ACCESS_SECRET"

// Redis commands
const SET = "SET"
const EXISTS = "EXISTS"
const GET = "GET"
const INCR = "INCR"

// Handler methods
const POST = "POST"

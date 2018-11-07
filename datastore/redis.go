package datastore

import (
	"errors"
	"fmt"
	"log"
	"twittertracker/common"
	"twittertracker/models"

	"github.com/mediocregopher/radix.v2/pool"
)

type RedisDatastore struct {
	*pool.Pool
}

// NewRedisDatastore initializes new thread pool connected to Redis
func NewRedisDatastore(address string) (*RedisDatastore, error) {

	connectionPool, err := pool.New(common.RedisQueueProtocol, address, 10)
	if err != nil {
		return nil, err
	}
	return &RedisDatastore{
		Pool: connectionPool,
	}, nil
}

// CreateExample is an example of how to set key values in Redis
func (r *RedisDatastore) CreateExample(example *models.Example) error {

	if r.Cmd(common.Set, example.ExampleId, example.ExampleValue).Err != nil {
		return errors.New("Failed to execute Redis SET command")
	}

	return nil
}

// GetExample is an example of how to retrieve the value of a key in Redis
func (r *RedisDatastore) GetExample(exampleID string) (*models.Example, error) {

	exists, err := r.Cmd(common.Exists, exampleID).Int()

	if err != nil {
		return nil, err
	} else if exists == 0 {
		return nil, nil
	}

	exampleVal, err := r.Cmd(common.Get, exampleID).Str()
	fmt.Println(exampleVal)

	if err != nil {
		log.Print(err)

		return nil, err
	}

	return &models.Example{ExampleId: exampleID, ExampleValue: exampleVal}, nil
}

// SaveWord saves word into Redis
func (r *RedisDatastore) SaveWord(word string) error {
	if r.Cmd(common.Incr, word).Err != nil {
		return errors.New("Failed to increment key" + word + "by 1")
	}

	return nil
}

// BeginTransaction initializes the transaction for Redis
func (r *RedisDatastore) BeginTransaction() error {
	if r.Cmd(common.Multi).Err != nil {
		return errors.New("Failed to begin transaction")
	}
	return nil
}

// ExecTransaction executes the transaction for redis
func (r *RedisDatastore) ExecTransaction() error {
	if r.Cmd(common.Exec).Err != nil {
		return errors.New("Failed to execute transaction")
	}
	return nil
}

// Close closes the db connection
func (r *RedisDatastore) Close() {
	r.Empty()
}

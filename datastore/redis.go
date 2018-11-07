package datastore

import (
	"errors"
	"fmt"
	"log"
	"twittertracker/models"

	"github.com/mediocregopher/radix.v2/pool"
)

type RedisDatastore struct {
	*pool.Pool
}

func NewRedisDatastore(address string) (*RedisDatastore, error) {

	connectionPool, err := pool.New("tcp", address, 10)
	if err != nil {
		return nil, err
	}
	return &RedisDatastore{
		Pool: connectionPool,
	}, nil
}

func (r *RedisDatastore) CreateExample(example *models.Example) error {

	if r.Cmd("SET", example.ExampleId, example.ExampleValue).Err != nil {
		return errors.New("Failed to execute Redis SET command")
	}

	return nil
}

func (r *RedisDatastore) GetExample(exampleId string) (*models.Example, error) {

	exists, err := r.Cmd("EXISTS", exampleId).Int()

	if err != nil {
		return nil, err
	} else if exists == 0 {
		return nil, nil
	}

	exampleVal, err := r.Cmd("GET", exampleId).Str()
	fmt.Println(exampleVal)

	if err != nil {
		log.Print(err)

		return nil, err
	}

	return &models.Example{ExampleId: exampleId, ExampleValue: exampleVal}, nil
}

func (r *RedisDatastore) SaveWord(word string) error {
	if r.Cmd("INCR", word).Err != nil {
		return errors.New("Failed to increment key" + word + "by 1")
	}

	return nil
}

func (r *RedisDatastore) Close() {
	r.Empty()
}

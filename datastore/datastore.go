package datastore

import (
	"errors"
	"twittertracker/models"
)

type Datastore interface {
	CreateExample(user *models.Example) error
	GetExample(username string) (*models.Example, error)
	Close()
}

const (
	REDIS = iota
	// MONGODB
	// MYSQL
)

func NewDatastore(datastoreType int, dbConnectionString string) (Datastore, error) {

	switch datastoreType {
	case REDIS:
		return NewRedisDatastore(dbConnectionString)
	default:
		return nil, errors.New("The datastore you specified does not exist!")
	}
}

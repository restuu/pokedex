package mgo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DbName is an alias of string to accomodate usage of google wire binding
type DbName string

func (d DbName) String() string {
	return string(d)
}

func NewDatabase(client *mongo.Client, dbName DbName) *mongo.Database {
	return client.Database(dbName.String(), options.Database())
}

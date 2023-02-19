package mgo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// DbUri is an alias of string to accomodate usage of google wire binding
type DbUri string

func (d DbUri) String() string {
	return string(d)
}

func NewClient(ctx context.Context, uri DbUri) (*mongo.Client, error) {
	opt := options.Client().ApplyURI(uri.String())

	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		return nil, err
	}

	// ping
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client, nil
}

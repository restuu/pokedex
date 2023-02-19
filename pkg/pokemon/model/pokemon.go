package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Pokemon struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string             `bson:"name" json:"name"`
}

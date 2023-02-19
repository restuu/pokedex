package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type CapturedPokemon struct {
	ID         primitive.ObjectID   `bson:"_id"`
	UserID     primitive.ObjectID   `bson:"user_id"`
	PokemonIDs []primitive.ObjectID `bson:"pokemon_ids"`
}

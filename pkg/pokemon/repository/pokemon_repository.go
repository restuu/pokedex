package repository

import (
	"context"
	"pokedex/pkg/pokemon/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PokemonRepository interface {
	Add(ctx context.Context, pokemon model.Pokemon) (*primitive.ObjectID, error)
}

type pokemonMgoRepo struct {
	col *mongo.Collection
}

const (
	collectionName = "pokemons"
)

func NewPokemonRepository(db *mongo.Database) PokemonRepository {
	return &pokemonMgoRepo{
		col: db.Collection(collectionName),
	}
}

func (r *pokemonMgoRepo) Add(ctx context.Context, pokemon model.Pokemon) (*primitive.ObjectID, error) {
	result, err := r.col.InsertOne(ctx, pokemon)
	if err != nil {
		return nil, err
	}

	oid := result.InsertedID.(primitive.ObjectID)
	return &oid, nil
}

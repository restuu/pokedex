package repository

import (
	"context"
	"pokedex/pkg/pokemon/model"

	"gorm.io/gorm"
)

type PokemonRepository interface {
	Add(ctx context.Context, pokemon model.Pokemon) (*uint, error)
}

type pokemonMgoRepo struct {
	model *gorm.DB
}

func NewPokemonRepository(db *gorm.DB) PokemonRepository {

	db.AutoMigrate(model.Pokemon{})

	return &pokemonMgoRepo{
		model: db.Model(model.Pokemon{}),
	}
}

func (r *pokemonMgoRepo) Add(ctx context.Context, pokemon model.Pokemon) (*uint, error) {
	err := r.model.Create(&pokemon).Error

	if err != nil {
		return nil, err
	}

	return &pokemon.ID, nil
}

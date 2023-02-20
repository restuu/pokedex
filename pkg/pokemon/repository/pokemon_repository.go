package repository

import (
	"context"
	"fmt"
	"pokedex/pkg/pokemon/model"
	"sync"

	"gorm.io/gorm"
)

type PokemonRepository interface {
	Add(ctx context.Context, pokemon model.Pokemon) (*model.Pokemon, error)
}

type pokemonMgoRepo struct {
	model *gorm.DB
}

var (
	onceRepo     = sync.Once{}
	_pokemonRepo PokemonRepository
)

func NewPokemonRepository(db *gorm.DB) PokemonRepository {

	onceRepo.Do(func() {

		db.AutoMigrate(model.Pokemon{})

		_pokemonRepo = &pokemonMgoRepo{
			model: db.Model(model.Pokemon{}),
		}
	})

	return _pokemonRepo
}

func (r *pokemonMgoRepo) Add(ctx context.Context, pokemon model.Pokemon) (*model.Pokemon, error) {

	newPokemon := new(model.Pokemon)

	err := r.model.
		Where(model.Pokemon{Name: pokemon.Name}).
		Attrs(pokemon).
		FirstOrCreate(newPokemon).
		Error

	fmt.Println(newPokemon)

	return newPokemon, err
}

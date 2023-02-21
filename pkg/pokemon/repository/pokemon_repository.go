package repository

import (
	"context"
	"errors"
	"fmt"
	"pokedex/pkg/pokemon/model"
	"sync"

	"gorm.io/gorm"
)

type PokemonRepository interface {
	Add(ctx context.Context, pokemon model.Pokemon) (*model.Pokemon, error)
	FindAll(ctx context.Context) ([]model.Pokemon, error)
	FindByID(ctx context.Context, id uint) (*model.Pokemon, error)
	Save(ctx context.Context, pokemon model.Pokemon) (*model.Pokemon, error)
}

type pokemonRepo struct {
	model *gorm.DB
}

var (
	onceRepo     = sync.Once{}
	_pokemonRepo PokemonRepository
)

func NewPokemonRepository(db *gorm.DB) PokemonRepository {

	onceRepo.Do(func() {

		db.AutoMigrate(model.Pokemon{})

		_pokemonRepo = &pokemonRepo{
			model: db.Model(model.Pokemon{}),
		}
	})

	return _pokemonRepo
}

func (r *pokemonRepo) Add(ctx context.Context, pokemon model.Pokemon) (*model.Pokemon, error) {

	newPokemon := new(model.Pokemon)

	err := r.model.
		Where(model.Pokemon{Name: pokemon.Name}).
		Attrs(pokemon).
		FirstOrCreate(newPokemon).
		WithContext(ctx).
		Error

	fmt.Println(newPokemon)

	return newPokemon, err
}

func (r *pokemonRepo) FindAll(ctx context.Context) ([]model.Pokemon, error) {

	all := []model.Pokemon{}

	err := r.model.
		WithContext(ctx).
		Find(&all).
		Error

	return all, err
}

func (r *pokemonRepo) FindByID(ctx context.Context, id uint) (*model.Pokemon, error) {
	m := new(model.Pokemon)

	err := r.model.
		First(m, model.Pokemon{ID: id}).
		WithContext(ctx).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return m, nil
}

func (r *pokemonRepo) Save(ctx context.Context, pokemon model.Pokemon) (*model.Pokemon, error) {

	updated := pokemon

	err := r.model.
		WithContext(ctx).
		Save(&updated).
		Error

	return &updated, err
}

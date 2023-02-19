package service

import (
	"context"
	"pokedex/pkg/pokemon/model"
	"pokedex/pkg/pokemon/repository"
)

type PokemonAddingService interface {
	Add(ctx context.Context, pokemon model.Pokemon) (*model.Pokemon, error)
}

func NewPokemonAddingService(
	pokemonRepository repository.PokemonRepository,

) PokemonAddingService {

	return &pokemonAddingService{
		pokemonRepo: pokemonRepository,
	}
}

type pokemonAddingService struct {
	pokemonRepo repository.PokemonRepository
}

func (s *pokemonAddingService) Add(ctx context.Context, pokemon model.Pokemon) (*model.Pokemon, error) {
	newID, err := s.pokemonRepo.Add(ctx, pokemon)
	if err != nil {
		return nil, err
	}

	pokemon.ID = *newID

	return &pokemon, nil
}

package service

import (
	"context"
	"pokedex/pkg/myerror"
	"pokedex/pkg/pokemon/dto"
	"pokedex/pkg/pokemon/repository"
	"pokedex/pkg/pokemon/util"
	"sync"
)

type PokemonUpdatingService interface {
	UpdateOne(ctx context.Context, pokemon dto.PokemonUpdateRequest) (*dto.PokemonResponse, error)
}

var (
	_syncUpdatingService    = sync.Once{}
	_pokemonUpdatingService PokemonUpdatingService
)

func NewPokemonUpdatingService(
	pokemonRepo repository.PokemonRepository,

) PokemonUpdatingService {

	_syncUpdatingService.Do(func() {
		_pokemonUpdatingService = &pokemonUpdatingService{
			pokemonRepo: pokemonRepo,
		}
	})

	return _pokemonUpdatingService
}

type pokemonUpdatingService struct {
	pokemonRepo repository.PokemonRepository
}

func (s *pokemonUpdatingService) UpdateOne(ctx context.Context, pokemon dto.PokemonUpdateRequest) (*dto.PokemonResponse, error) {
	if pokemon.ID == 0 {
		return nil, myerror.ErrInvalidFormat
	}

	existingPokemon, err := s.pokemonRepo.FindByID(ctx, pokemon.ID)
	if err != nil {
		return nil, err
	}
	if existingPokemon == nil {
		return nil, myerror.ErrDataNotFound
	}

	existingPokemon.Name = pokemon.Name
	existingPokemon.ImageURL = pokemon.ImageURL
	existingPokemon.Type = pokemon.Type

	updatedPokemon, err := s.pokemonRepo.Save(ctx, *existingPokemon)
	if err != nil {
		return nil, err
	}

	response := util.PokemonToPokemonResponse(*updatedPokemon)
	return &response, nil
}

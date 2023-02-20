package service

import (
	"context"
	"pokedex/pkg/pokemon/dto"
	"pokedex/pkg/pokemon/repository"
	"pokedex/pkg/pokemon/util"
	"sync"
)

type PokemonGettingService interface {
	GetAll(ctx context.Context) ([]dto.PokemonResponse, error)
}

var (
	onceGettingSrv         = sync.Once{}
	_pokemonGettingService PokemonGettingService
)

func NewPokemonGettingService(
	pokemonRepository repository.PokemonRepository,

) PokemonGettingService {
	onceGettingSrv.Do(func() {
		_pokemonGettingService = &pokemonGettingService{
			pokemonRepo: pokemonRepository,
		}
	})

	return _pokemonGettingService
}

type pokemonGettingService struct {
	pokemonRepo repository.PokemonRepository
}

func (s *pokemonGettingService) GetAll(ctx context.Context) ([]dto.PokemonResponse, error) {
	pokemons, err := s.pokemonRepo.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	pokemonResponses := make([]dto.PokemonResponse, len(pokemons))
	for i, v := range pokemons {
		pokemonResponses[i] = util.PokemonToPokemonResponse(v)
	}

	return pokemonResponses, nil
}

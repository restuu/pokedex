package service

import (
	"context"
	"pokedex/pkg/pokemon/dto"
	"pokedex/pkg/pokemon/model"
	"pokedex/pkg/pokemon/repository"
	"pokedex/pkg/pokemon/util"
)

type PokemonAddingService interface {
	Add(ctx context.Context, pokemon dto.PokemonAddRequest) (*dto.PokemonResponse, error)
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

func (s *pokemonAddingService) Add(ctx context.Context, pokemon dto.PokemonAddRequest) (*dto.PokemonResponse, error) {

	pokemonModel := model.Pokemon{
		Name: pokemon.Name,
	}

	newPokemon, err := s.pokemonRepo.Add(ctx, pokemonModel)
	if err != nil {
		return nil, err
	}

	pokemonResponse := util.PokemonToPokemonResponse(*newPokemon)
	return &pokemonResponse, nil

}

package service

import (
	"context"
	"pokedex/pkg/pokemon/dto"
	"pokedex/pkg/pokemon/model"
	"pokedex/pkg/pokemon/repository"
)

type PokemonAddingService interface {
	Add(ctx context.Context, pokemon dto.PokemonAddRequest) (*model.Pokemon, error)
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

func (s *pokemonAddingService) Add(ctx context.Context, pokemon dto.PokemonAddRequest) (*model.Pokemon, error) {

	pokemonModel := model.Pokemon{
		Name: pokemon.Name,
	}

	newID, err := s.pokemonRepo.Add(ctx, pokemonModel)
	if err != nil {
		return nil, err
	}

	pokemonModel.ID = *newID

	return &pokemonModel, nil
}

package util

import (
	"pokedex/pkg/pokemon/dto"
	"pokedex/pkg/pokemon/model"
)

func PokemonToPokemonResponse(pokemon model.Pokemon) dto.PokemonResponse {
	dto := dto.PokemonResponse{
		ID:        pokemon.ID,
		Name:      pokemon.Name,
		Type:      pokemon.Type,
		ImageURL:  pokemon.ImageURL,
		CreatedAt: pokemon.CreatedAt,
		UpdatedAt: pokemon.UpdatedAt,
	}

	if pokemon.DeletedAt.Valid {
		dto.DeletedAt = &pokemon.DeletedAt.Time
	}

	return dto
}

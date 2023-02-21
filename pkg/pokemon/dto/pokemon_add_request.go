package dto

import "pokedex/pkg/pokemon/constant"

type PokemonAddRequest struct {
	Name     string               `json:"name"`
	ImageURL string               `json:"image_url"`
	Type     constant.PokemonType `json:"type"`
}

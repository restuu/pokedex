package dto

import (
	"pokedex/pkg/pokemon/constant"
	"time"
)

type PokemonResponse struct {
	ID        uint                 `json:"id"`
	Name      string               `json:"name"`
	Type      constant.PokemonType `json:"type"`
	ImageURL  string               `json:"image_url"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
	DeletedAt *time.Time           `json:"deleted_at"`
}

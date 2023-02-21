package dto

type PokemonUpdateRequest struct {
	ID uint `json:"id"`

	PokemonAddRequest
}

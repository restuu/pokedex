package router

import (
	"net/http"
	"pokedex/pkg/pokemon/service"

	"github.com/labstack/echo/v4"
)

func NewPokemonRouter(e *echo.Echo,
	pokemonGettingService service.PokemonGettingService) {

	g := e.Group("/pokemons")

	h := &pokemonHandler{
		pokemonGettingService: pokemonGettingService,
	}

	g.GET("", h.findAll)
}

type pokemonHandler struct {
	pokemonGettingService service.PokemonGettingService
}

func (h *pokemonHandler) findAll(c echo.Context) error {
	pokemons, err := h.pokemonGettingService.GetAll(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, pokemons)
}

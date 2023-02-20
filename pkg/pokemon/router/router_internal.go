package router

import (
	"fmt"
	"pokedex/pkg/app"
	"pokedex/pkg/pokemon/dto"
	"pokedex/pkg/pokemon/service"

	"github.com/labstack/echo/v4"
)

func NewPokemonInternalRouter(e *echo.Echo,
	conf *app.Config,
	pokemonAddingService service.PokemonAddingService,
) {
	g := e.Group("/internal/pokemons")
	//TODO: Re enable when ready
	// g.Use(myjwt.Middleware(conf.JWTKey), myjwt.IsAdmin())

	h := &pokemonInternalHandler{
		pokemonAddingService: pokemonAddingService,
	}

	g.POST("", h.Add)

}

type pokemonInternalHandler struct {
	pokemonAddingService service.PokemonAddingService
}

func (h *pokemonInternalHandler) Add(c echo.Context) error {
	p := dto.PokemonAddRequest{}
	err := c.Bind(&p)
	if err != nil {
		return err
	}

	newPokemon, err := h.pokemonAddingService.Add(c.Request().Context(), p)
	if err != nil {
		return err
	}

	fmt.Println("--------------")
	fmt.Println(newPokemon)

	c.JSON(200, newPokemon)

	return nil
}

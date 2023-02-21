package router

import (
	"pokedex/pkg/app"
	"pokedex/pkg/pokemon/dto"
	"pokedex/pkg/pokemon/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

func NewPokemonInternalRouter(e *echo.Echo,
	conf *app.Config,
	pokemonAddingService service.PokemonAddingService,
	pokemonUpdatingService service.PokemonUpdatingService,
) {
	g := e.Group("/internal/pokemons")
	//TODO: Re enable when ready
	// g.Use(myjwt.Middleware(conf.JWTKey), myjwt.IsAdmin())

	h := &pokemonInternalHandler{
		pokemonAddingService:   pokemonAddingService,
		pokemonUpdatingService: pokemonUpdatingService,
	}

	g.POST("", h.Add)
	g.PUT("/:id", h.updateOne)

}

type pokemonInternalHandler struct {
	pokemonAddingService   service.PokemonAddingService
	pokemonUpdatingService service.PokemonUpdatingService
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

	c.JSON(200, newPokemon)

	return nil
}

func (h *pokemonInternalHandler) updateOne(c echo.Context) error {
	req := new(dto.PokemonUpdateRequest)

	if err := c.Bind(req); err != nil {
		return err
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return err
	}

	req.ID = uint(id)

	r, err := h.pokemonUpdatingService.UpdateOne(c.Request().Context(), *req)
	if err != nil {
		return err
	}

	return c.JSON(200, r)

}

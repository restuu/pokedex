//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"pokedex/pkg/app"
	"pokedex/pkg/datastore/mysql"
	pokemonRepository "pokedex/pkg/pokemon/repository"
	pokemonService "pokedex/pkg/pokemon/service"

	"github.com/google/wire"
)

func initializeApp(ctx context.Context, conf *app.Config, dbUri string) (*server, error) {

	wire.Build(
		NewServer,
		wire.Struct(new(service), "*"),
		mysql.Connect,
		pokemonRepository.NewPokemonRepository,
		pokemonService.NewPokemonAddingService,
	)
	return nil, nil // This will be overwritten by the wire compiler.
}

type service struct {
	pokemonAddingService pokemonService.PokemonAddingService
}

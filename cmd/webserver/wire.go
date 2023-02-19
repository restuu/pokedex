//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"pokedex/pkg/datastore/mgo"
	pokemonRepository "pokedex/pkg/pokemon/repository"
	pokemonService "pokedex/pkg/pokemon/service"

	"github.com/google/wire"
)

func initializeApp(ctx context.Context, dbUri mgo.DbUri, dbName mgo.DbName) (*App, error) {

	wire.Build(NewApp,
		pokemonService.NewPokemonAddingService,
		pokemonRepository.NewPokemonRepository,
		mgo.NewDatabase,
		mgo.NewClient,
	)

	return &App{}, nil
}

type App struct {
	pokemonAddingService pokemonService.PokemonAddingService
}

func NewApp(
	pokemonAddingService pokemonService.PokemonAddingService,
) *App {
	return &App{
		pokemonAddingService: pokemonAddingService,
	}
}

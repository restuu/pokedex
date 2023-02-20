package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"pokedex/pkg/app"
	pokemonRouter "pokedex/pkg/pokemon/router"
)

func start(ctx context.Context, conf app.Config) {

	e, err := initializeApp(ctx, &conf, conf.DbUri)

	if err != nil {
		log.Fatal(err)
	}

	e.Start()
}

type server struct {
	conf    *app.Config
	service *service
}

func NewServer(s *service, conf *app.Config) *server {

	return &server{
		conf:    conf,
		service: s,
	}

}

func (s *server) Start() {

	baseCtx := context.Background()

	e := route()

	//#region register router
	pokemonRouter.NewPokemonInternalRouter(e, s.conf, s.service.pokemonAddingService)
	pokemonRouter.NewPokemonRouter(e, s.service.pokemonGettingService)
	//#endregion register router

	svr := http.Server{
		Addr:    fmt.Sprintf(":%d", s.conf.Port),
		Handler: e,
		BaseContext: func(l net.Listener) context.Context {
			return baseCtx
		},
	}

	log.Printf("listening to: %d\n", s.conf.Port)
	if err := svr.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

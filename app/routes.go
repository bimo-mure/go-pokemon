package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	mux := httprouter.New()

	mux.GET("/pokemon-list", app.PokemonList)
	mux.GET("/pokemon-detail/:name", app.PokemonDetail)

	return mux
}

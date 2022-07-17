package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	mux := httprouter.New()

	mux.GET("/pokemon-list", app.PokemonList)
	mux.GET("/pokemon-detail", app.PokemonDetail)

	return mux
}

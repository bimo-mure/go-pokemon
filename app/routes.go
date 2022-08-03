package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	mux := httprouter.New()

	mux.GET("/pokemon-list", app.PokemonList)
	mux.GET("/pokemon-detail/:name", app.PokemonDetail)
	mux.GET("/my-pokemon", app.MyPokemon)
	mux.GET("/get-prime-number", app.GetPrimeNumber)
	mux.POST("/save-pokemon", app.SavePokemon)
	mux.POST("/catch-pokemon", app.CatchPokemon)
	mux.DELETE("/release-pokemon", app.ReleasePokemon)

	return mux
}

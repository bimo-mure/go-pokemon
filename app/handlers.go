package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) PokemonList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	listPokemon, err := app.GetPokemonList(w, r, nil)

	err = app.renderTemplate(w, r, "pokemon", &templateData{
		ListPokemon: listPokemon,
	})
	if err != nil {
		app.errorLog.Println(err)
	}
}

func (app *application) PokemonDetail(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	listPokemon, err := app.GetPokemonList(w, r, nil)

	err = app.renderTemplate(w, r, "pokemon-detail", &templateData{
		ListPokemon: listPokemon,
	})
	if err != nil {
		app.errorLog.Println(err)
	}
}

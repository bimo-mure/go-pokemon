package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dev-bimomure/go-pokemon/app/structs"
	"github.com/julienschmidt/httprouter"
)

func (app *application) PokemonList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var result structs.Pokemon
	var list []structs.PokemonSource

	for i := 0; i < 10; i++ {
		pokemon, err := app.call(fmt.Sprintf("pokemon/%d", i+1))
		if err != nil {
			app.errorLog.Println(err)
		}
		json.Unmarshal(pokemon, &result)

		pokemonSource := structs.PokemonSource{
			Image: result.Sprites.Other.OfficialArtwork.FrontDefault,
			Name:  result.Name,
		}

		list = append(list, pokemonSource)
	}

	err := app.renderTemplate(w, r, "pokemon", &templateData{
		PokemonList: list,
	})
	if err != nil {
		app.errorLog.Println(err)
	}
}

func (app *application) PokemonDetail(w http.ResponseWriter, r *http.Request, data httprouter.Params) {
	var result structs.Pokemon
	pokemon, err := app.call(fmt.Sprintf("pokemon/%s", data.ByName("name")))
	if err != nil {
		app.errorLog.Println(err)
	}
	json.Unmarshal(pokemon, &result)

	var moves []string
	for _, item := range result.Moves {
		moves = append(moves, item.Move.Name)
	}

	var types []string
	for _, item := range result.Types {
		types = append(types, item.Type.Name)
	}

	pokemonDetail := structs.PokemonDetail{
		Image: result.Sprites.Other.OfficialArtwork.FrontDefault,
		Name:  result.Name,
		Moves: moves,
		Types: types,
	}

	err = app.renderTemplate(w, r, "pokemon.detail", &templateData{
		PokemonDetail: pokemonDetail,
	})
	if err != nil {
		app.errorLog.Println(err)
	}
}

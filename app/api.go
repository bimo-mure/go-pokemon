package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Pokemon struct {
	Name string
	Url  string
}

func (app *application) GetPokemonList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) ([]Pokemon, error) {
	var pokemonList []Pokemon
	pokemon := make(map[string]interface{})

	response, err := http.Get("https://pokeapi.co/api/v2/pokemon")
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		app.errorLog.Println(err)
		return nil, err
	}

	err = json.Unmarshal(responseBody, &pokemon)
	if err != nil {
		app.errorLog.Println(err)
		return nil, err
	}

	dataInterf := getResultInterface(pokemon)

	for _, item := range dataInterf {
		oke := item.(map[string]interface{})

		poke := getUrlNameInterface(oke)

		pokemonList = append(pokemonList, poke)
	}

	return pokemonList, nil
}

func getResultInterface(data interface{}) []interface{} {
	m := data.(map[string]interface{})
	var arrInter []interface{}
	if results, ok := m["results"].([]interface{}); ok {
		arrInter = results
	}

	return arrInter
}

func getUrlNameInterface(data interface{}) Pokemon {
	m := data.(map[string]interface{})
	var pokeData Pokemon
	if name, ok := m["name"].(string); ok {
		pokeData.Name = name
	}
	if url, ok := m["url"].(string); ok {
		pokeData.Url = url
	}

	return pokeData
}

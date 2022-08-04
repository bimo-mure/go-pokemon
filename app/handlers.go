package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dev-bimomure/go-pokemon/app/structs"
	"github.com/julienschmidt/httprouter"
)

type PokeResponse struct {
	Code    int
	Data    structs.MyPokemon
	Message string
}

type CatchResponse struct {
	Code    int
	Result  int
	Message string
}

type GetNumberResponse struct {
	Code    int
	Result  int
	IsPrime bool
}

type GeneralResponse struct {
	Code         int
	RowsAffected int
	Message      string
}

func (app *application) PokemonList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var result structs.Pokemon
	var list []structs.PokemonSource

	for i := 0; i < 9; i++ {
		randomPokemonId := GenerateRandomNumberPokemon()
		pokemon, err := app.callPokeAPI(fmt.Sprintf("pokemon/%d", randomPokemonId))
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
	pokemon, err := app.callPokeAPI(fmt.Sprintf("pokemon/%s", data.ByName("name")))
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
	stats := structs.PokemonStats{
		HP:             result.Stats[0].BaseStat,
		Attack:         result.Stats[1].BaseStat,
		Defense:        result.Stats[2].BaseStat,
		SpecialAttack:  result.Stats[3].BaseStat,
		SpecialDefense: result.Stats[4].BaseStat,
		Speed:          result.Stats[5].BaseStat,
	}

	pokemonDetail := structs.PokemonDetail{
		Image: result.Sprites.Other.OfficialArtwork.FrontDefault,
		Name:  result.Name,
		Moves: moves,
		Types: types,
		Stats: stats,
	}

	err = app.renderTemplate(w, r, "pokemon.detail", &templateData{
		PokemonDetail: pokemonDetail,
	})
	if err != nil {
		app.errorLog.Println(err)
	}
}

func (app *application) MyPokemon(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	myPokemonList, err := app.DB.GetPokemon()
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	err = app.renderTemplate(w, r, "my.pokemon", &templateData{
		MyPokemonList: myPokemonList,
	})
	if err != nil {
		app.errorLog.Println(err)
	}

}

func (app *application) SavePokemon(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	bodyReq, _ := ioutil.ReadAll(r.Body)
	var pokemon structs.MyPokemon
	var code int
	message := fmt.Sprintf("Successfully Insert Pokemon %s", pokemon.PokemonName)

	json.Unmarshal(bodyReq, &pokemon)
	_, err := app.DB.InsertPokemon(pokemon)
	if err != nil {
		app.errorLog.Println(err)
		code = 500
		message = "Internal Server error"
	}

	result := PokeResponse{
		Code:    code,
		Data:    pokemon,
		Message: message,
	}
	res, err := json.Marshal(result)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (app *application) CatchPokemon(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	bodyReq, _ := ioutil.ReadAll(r.Body)
	var pokemon structs.MyPokemon
	var result CatchResponse
	json.Unmarshal(bodyReq, &pokemon)
	coin := FlipCoin()

	if coin == 1 {
		result = CatchResponse{
			Code:    200,
			Result:  coin,
			Message: fmt.Sprintf("Successfully Catch Pokemon %s", pokemon.PokemonName),
		}
	} else {
		result = CatchResponse{
			Code:    200,
			Result:  coin,
			Message: fmt.Sprintf("Failed Catch Pokemon %s", pokemon.PokemonName),
		}
	}

	res, err := json.Marshal(result)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (app *application) callPokeAPI(endpoint string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, app.config.apiurl+endpoint, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}

func (app *application) GetPrimeNumber(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	RandNumber := GenerateRandomNumber()
	isPrime := CheckPrimeNumber(RandNumber)

	result := GetNumberResponse{
		Code:    200,
		Result:  RandNumber,
		IsPrime: isPrime,
	}

	res, err := json.Marshal(result)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (app *application) ReleasePokemon(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	bodyReq, _ := ioutil.ReadAll(r.Body)
	var pokemon structs.MyPokemon
	json.Unmarshal(bodyReq, &pokemon)
	rowsAffected, err := app.DB.DeletePokemon(pokemon.Id)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	result := GeneralResponse{
		Code:         200,
		RowsAffected: rowsAffected,
		Message:      fmt.Sprintf("Successfully Delete Pokemon %s", pokemon.PokemonName),
	}
	res, err := json.Marshal(result)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (app *application) RenamePokemon(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	bodyReq, _ := ioutil.ReadAll(r.Body)
	var pokemon structs.MyPokemon
	json.Unmarshal(bodyReq, &pokemon)

	var newNickname string
	var countUpdate int

	lastUpdate, err := app.DB.GetPokemonLastCountById(pokemon.Id)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	if lastUpdate > 0 {
		countUpdate = lastUpdate + 1
		fib := Fibonacci(countUpdate)
		newNickname = fmt.Sprintf("%s-%d", pokemon.NickName, fib)
	} else {
		newNickname = fmt.Sprintf("%s-%d", pokemon.NickName, 0)
		countUpdate = 1
	}

	rowsAffected, err := app.DB.UpdatePokemon(pokemon.Id, newNickname, countUpdate)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	result := GeneralResponse{
		Code:         200,
		RowsAffected: rowsAffected,
		Message:      fmt.Sprintf("Successfully Update nickname %s with id %d", newNickname, pokemon.Id),
	}
	res, err := json.Marshal(result)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

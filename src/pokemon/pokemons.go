package pokemon

import (
	"net/http"
	"encoding/json"
	"util"
)

var pokemons []Pokemon

func AddPokemon(w http.ResponseWriter, r *http.Request) {
	pokemons.Append(&Pokemon{
		Id: 0,
		Name: r.URL.Path[1:],
		Type: "truc",
		Description: "awesome desc",
	})

	res, err := json.Marshal(pokemons)

	util.RespCreated(w, res, err)
}
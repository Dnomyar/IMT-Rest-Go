package pokemon

import (
	"net/http"
	"encoding/json"
	"fmt"
)

var pokemons []Pokemon = []Pokemon{

	{
		Id: 25,
		Name: "Pikachu",
		Type: "Foudre",
		Description: "La souris électrique",
	},
	{
		Id: 25,
		Name: "Pikachu",
		Type: "Foudre",
		Description: "La souris électrique",
	},
	{
		Id: 4,
		Name: "Salamèche",
		Type: "Feu",
		Description: "La salamandre enflammée",
	},
	{
		Id: 7,
		Name: "Bulbizarre",
		Type: "Plante",
		Description: "Le machin avec un bulbe",
	},
	{
		Id: 3,
		Name: "Carapuce",
		Type: "Eau",
		Description: "La tortue qui te crache dessus",
	},
}

/**
pokemons = append(pokemons, Pokemon{
		Id: 25,
		Name: "Pikachu",
		Type: "Foudre",
		Description: "La souris électrique",
	})
 */

func AddPokemon(w http.ResponseWriter, r *http.Request) {

}

func GetPokemon(w http.ResponseWriter, r *http.Request) {
	res, err := json.Marshal(pokemons)
	util.RespJson(w, 200, res, err)
}

func AddPokemon(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%+v\n", r.Body)
	res, err := json.Marshal(pokemons)
	util.RespJson(w, 200, res, err)
}
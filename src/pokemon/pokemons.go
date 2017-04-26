package pokemon

import (
	"net/http"
	"encoding/json"
	"util"
	"log"
	"strconv"
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



func GetPokemons(w http.ResponseWriter, r *http.Request) {
	res, err := json.Marshal(pokemons)
	util.RespJson(w, 200, res, err)
}

func GetPokemon(w http.ResponseWriter, r *http.Request) {

	var id, err = strconv.ParseInt(r.URL.Path[1:], 10, 32)

	log.Println(err)

	res, err := json.Marshal(pokemons[id])

	log.Println(err)

	util.RespJson(w, 200, res, err)
}


func AddPokemon(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var pokemon Pokemon
	err := decoder.Decode(&pokemon)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	log.Println(pokemon)

	pokemons = append(pokemons, pokemon)

	res, err := json.Marshal(pokemons)
	util.RespJson(w, 200, res, err)
}

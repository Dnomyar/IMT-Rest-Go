package pokemon

import (
	"net/http"
	"encoding/json"
	"util"
)

var pokemons [150]Pokemon

func AddPokemon(w http.ResponseWriter, r *http.Request) {
	pokemons[24] = Pokemon{
		Id: 25,
		Name: "Pikachu",
		Type: "Foudre",
		Description: "La souris électrique",
	}
	pokemons[3] = Pokemon{
		Id: 4,
		Name: "Salamèche",
		Type: "Feu",
		Description: "La salamandre enflammée",
	}
	pokemons[6] = Pokemon{
		Id: 7,
		Name: "Bulbizarre",
		Type: "Plante",
		Description: "Le machin avec un bulbe",
	}
	pokemons[3] = Pokemon{
		Id: 3,
		Name: "Carapuce",
		Type: "Eau",
		Description: "La tortue qui te crache dessus",
	}

	res, err := json.Marshal(pokemons)

	util.RespJson(w, 200, res, err)
}
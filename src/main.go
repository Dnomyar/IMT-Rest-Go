package main

import "net/http"
import "util"
import (
	"pokemon"
	"util/filter"
	"fmt"
)

func main() {
	fmt.Println("Running server on port 8090")
	var filters = []filter.Filter{
		filter.FilterAuth{},
	}

	util.SetFilters(filters)

	util.Router(http.MethodGet, "/", pokemon.GetPokemon)
	util.Router(http.MethodGet, "/", pokemon.GetPokemons)
	util.Router(http.MethodPost, "/", pokemon.AddPokemon)
	util.Run("8090")
}


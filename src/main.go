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

	util.Router(http.MethodGet, "/", pokemon.AddPokemon)
	http.ListenAndServe(":8090", nil)
}


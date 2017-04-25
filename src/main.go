package main

import "net/http"
import "util"
import "pokemon"

func main() {
	util.Router(http.MethodGet, "/", pokemon.AddPokemon)
	http.ListenAndServe(":8090", nil)
}


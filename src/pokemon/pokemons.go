package pokemon

import (
	"net/http"
	"encoding/json"
	"util"
)

func AddPokemon(w http.ResponseWriter, r *http.Request) {
	test := &Pokemon{
		Id: 0,
		Name: r.URL.Path[1:],
		Type: "truc",
		Description: "awesome desc",
	}

	res, err := json.Marshal(test)

	util.RespCreated(w, res, err)
}
package util

import (
	"net/http"
	"errors"
	"util/filter"
)

var filters = []filter.Filter{}

func SetFilters(fs []filter.Filter) {
	filters = fs
}

func Router(code string, path string, handler func(http.ResponseWriter, *http.Request)) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if (code == r.Method) {
			for i := range filters {
				if(! filters[i].Filter(r)){
					RespError(w, errors.New("Error in filter"), http.StatusBadRequest)
					return
				}

			}
			handler(w, r)
		} else {
			RespError(w, errors.New("`" + r.Method + " " + path + "` does not exists!"), http.StatusNotFound)
		}
	})
}

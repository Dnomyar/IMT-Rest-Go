package util

import (
	"net/http"
	"errors"
	"util/filter"
	"fmt"
)

var filters = []filter.Filter{}

func SetFilters(fs []filter.Filter) {
	filters = fs
}

func Router(code string, path string, handler func(http.ResponseWriter, *http.Request)) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if (code == r.Method) {
			for i := range filters {
				fmt.Println(filters[i].Filter(r))
			}
			handler(w, r)
		} else {
			RespError(w, errors.New("`" + r.Method + " " + path + "` does not exists!"), http.StatusNotFound)
		}
	})
}

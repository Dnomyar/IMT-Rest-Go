package util

import (
	"net/http"
	"errors"
)

func Router(code string, path string, handler func(http.ResponseWriter, *http.Request)) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if (code == r.Method) {
			handler(w, r)
		} else {
			RespError(w, errors.New("`" + r.Method + " " + path + "` does not exists!"), http.StatusNotFound)
		}
	})
}

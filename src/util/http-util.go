package util

import "net/http"

func RespError(w http.ResponseWriter, err error, status int) {
	http.Error(w, err.Error(), status)
}

func RespJson(w http.ResponseWriter, code int, body []byte, err error) {

	if err != nil {
		RespError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(body)
}



func RespCreated(w http.ResponseWriter, body []byte, err error) {
	RespJson(w, http.StatusCreated, body, err)
}


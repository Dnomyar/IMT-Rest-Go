package filter

import "net/http"

type Filter interface {
	Filter(r *http.Request) bool
}


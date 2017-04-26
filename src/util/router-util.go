package util

import (
	"net/http"
	"errors"
	"util/filter"
)

type request struct {
	Path string
	Method map[string]func(http.ResponseWriter, *http.Request)
}

var requests = []request{}

var filters = []filter.Filter{}

func SetFilters(fs []filter.Filter) {
	filters = fs
}


func Router(code string, path string, handler func(http.ResponseWriter, *http.Request)) {


	for i := range requests {
		var req = requests[i]
		if (req.Path == path) {
			req.Method[code] = handler
			return
		}

	}

	var emptyMap = make(map[string]func(http.ResponseWriter, *http.Request))
	emptyMap[code] = handler

	requests = append(requests, request{
		Path: path,
		Method: emptyMap,
	})

}


func Run(port string) {

	for i := range requests {
		var req = requests[i]

		http.HandleFunc(req.Path, func(w http.ResponseWriter, r *http.Request) {

			var isMethodFound = false

			for code := range req.Method {
				if (code == r.Method) {
					for i := range filters {
						if(! filters[i].Filter(r)){
							RespError(w, errors.New("Error in filter"), http.StatusBadRequest)
							return
						}

					}
					req.Method[code](w, r)
					isMethodFound = true
				}
			}

			if(!isMethodFound){
				RespError(w, errors.New("`" + r.Method + " " + req.Path + "` does not exists!"), http.StatusNotFound)
			}

		})
	}

	http.ListenAndServe(":" + port, nil)


}
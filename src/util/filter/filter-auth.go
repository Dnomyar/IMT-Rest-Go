package filter

import (
	"fmt"
	"net/http"
)

type FilterAuth struct {}


func (f FilterAuth) Filter(r *http.Request) bool {

	var authorisationHeader = r.Header.Get("Authorization")

	fmt.Println("Autorisation header : " + authorisationHeader)
	return authorisationHeader == "1234"
}
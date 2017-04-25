package filter

import (
	"fmt"
	"net/http"
)

type FilterAuth struct {}


func (f FilterAuth) Filter(r *http.Request) bool {
	fmt.Println("test")
	return false
}
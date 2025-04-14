package api

import (
	"fmt"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	text := "Hello"

	fmt.Fprint(w, text)
}

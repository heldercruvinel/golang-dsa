package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/heldercruvinel/golang-dsa/api"
)

type apiHandler struct{}

func (apiHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	text := "API"

	fmt.Fprint(w, text)
}

func main() {
	port := 8080
	apiHandler := apiHandler{}

	http.Handle("/api/", apiHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	http.HandleFunc("GET /test", api.Test)

	slog.Info(fmt.Sprintf("Server running on port %d", port))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		os.Exit(1)
	}
}

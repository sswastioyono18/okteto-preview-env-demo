package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	fmt.Println("This is okteto print test")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world this is endpoint /"))
	})

	r.Get("/v1", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome v1"))
	})

	http.ListenAndServe(":8080", r)
}
package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	fmt.Println("This is v4 tests")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("halo dunia tipu tipu"))
	})

	r.Get("/v1", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome v4"))
	})

	http.ListenAndServe(":8080", r)
}
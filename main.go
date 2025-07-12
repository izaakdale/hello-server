package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/izaakdale/ittp"
	"github.com/rs/cors"
)

func main() {
	mux := ittp.NewServeMux()

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"hello": "world"})
	})
	mux.Get("/{name}", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"hello": r.PathValue("name")})
	})

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	http.ListenAndServe(":"+port, c.Handler(mux))
}

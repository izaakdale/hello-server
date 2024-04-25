package main

import (
	"net/http"
	"os"

	"github.com/izaakdale/ittp"
	"github.com/rs/cors"
)

func main() {
	mux := ittp.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	mux.HandleFunc("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, " + r.PathValue("name") + "!"))
	})

	cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	http.ListenAndServe(":"+port, mux)
}

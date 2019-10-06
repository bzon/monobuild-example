package main

import (
	"log"
	"net/http"

	"github.com/bzon/monobuild-example/internal"
	"github.com/bzon/monobuild-example/pkg/api"
)

func main() {
	mux := internal.NewMux()
	mux.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		say := r.URL.Query().Get("say")
		w.Write([]byte(api.Echo(say)))
	})
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

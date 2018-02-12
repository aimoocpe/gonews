package main

import (
	"net/http"

	"github.com/aimoocpe/gonews/pkg/app"
)

const port = ":8000"

func main() {
	mux := http.NewServeMux()
	app.Mount(mux)
	http.ListenAndServe(port, mux)
}

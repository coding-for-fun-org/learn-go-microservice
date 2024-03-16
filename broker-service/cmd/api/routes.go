package main

import "net/http"

func (app *Config) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /broker", app.Broker)

	return mux
}

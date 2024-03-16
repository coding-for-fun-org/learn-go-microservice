package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Server running on port %s\n", webPort)

	// define http Server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start the Server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

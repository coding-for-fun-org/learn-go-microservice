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

	log.Printf("Broker server running on port %s\n", webPort)
}

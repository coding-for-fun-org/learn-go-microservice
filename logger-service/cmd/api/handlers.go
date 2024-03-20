package main

import (
	"log"
	"logger/data"
	"net/http"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	log.Print("Hit the /log endpoint")

	var reqeustPayload JSONPayload
	_ = app.readJSON(w, r, &reqeustPayload)

	event := data.LogEntry{
		Name: reqeustPayload.Name,
		Data: reqeustPayload.Data,
	}

	err := app.Models.LogEntry.Insert(event)
	if err != nil {
		app.errorJSON(w, err)

		return
	}

	response := jsonResponse{
		Error:   false,
		Message: "logged",
	}

	app.writeJSON(w, http.StatusAccepted, response)
}

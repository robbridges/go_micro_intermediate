package main

import (
	"log-service/data"
	"net/http"
)

type JSONpayload struct {
	Name string `json:name`
	Data string `json:"data"`
}

func (app *App) WriteLog(w http.ResponseWriter, r *http.Request) {
	var requestPayload JSONpayload

	_ = app.readJSON(w, r, &requestPayload)

	event := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}

	err := app.Models.LogEntry.Insert(event)

	if err != nil {
		app.errorJSON(w, err)
	}

	resp := jsonResponse{
		Error:   false,
		Message: "logged",
	}

	app.writeJSON(w, http.StatusAccepted, resp)
}

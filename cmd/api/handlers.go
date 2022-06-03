package main

import (
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := JsonResponse{
		Error:   false,
		Message: "Hit the broker endpoint",
	}

	err := app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		app.writeError(w, err)
	}

}

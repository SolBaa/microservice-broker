package main

import (
	"log"
	"net/http"
)

const webPORT = "8080"

type Config struct{}

func main() {
	app := Config{}
	log.Println("Starting broker service on port " + webPORT)

	//Define http server
	srv := &http.Server{
		Addr:    ":" + webPORT,
		Handler: app.routes(),
	}
	//Start http server

	log.Fatal(srv.ListenAndServe())
}

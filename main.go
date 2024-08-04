package main

import (
	"log"
	"net/http"

	"github.com/elliot-gustafsson03/corridor-server/api"
)

func main() {
	openServer()
}

func openServer() {
	api.AddSlideshowEndpoints()

	http.Handle("/apps/", http.StripPrefix("/apps", http.FileServer(http.Dir("apps"))))
	http.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("images"))))
	http.Handle("/", http.FileServer(http.Dir("public")))

	port := "7100"
	log.Print("Listening on port " + port + "...")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

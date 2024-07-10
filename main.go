package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var images ([2]Image)
var index = 0

func main() {
	images = loadImages()
	openServer()
}

type Image struct {
	Name  string `json:"image"`
	Label string `json:"label"`
}

func loadImages() [2]Image {
	array := [2]Image{{Name: "1.png", Label: "Bandpass-filter"}, {Name: "2.png", Label: "OP-amp"}}
	return array
}

func openServer() {
	fs := http.FileServer(http.Dir("./public"))

	http.Handle("/", fs)
	http.HandleFunc("/api/get_next_image", getNextImage)
	http.HandleFunc("/api/upload_image", uploadImage)

	port := "3333"

	log.Print("Listening on port " + port + "...")

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("err")
	}
}

// endpoints

func getNextImage(w http.ResponseWriter, r *http.Request) {

	json, err := json.Marshal(images[index%2])
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(500)
	}

	index++

	io.WriteString(w, string(json))
}

func uploadImage(w http.ResponseWriter, r *http.Request) {
	log.Println("tar emot bild")
}

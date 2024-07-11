package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/elliot-gustafsson03/corridor-server/models"
)

var images = models.List{}

func main() {
	openServer()
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

func getNextImage(w http.ResponseWriter, r *http.Request) {
	image := models.Image{Name: "#", Label: "No images uploaded yet"}

	if !images.IsEmpty() {
		image = *images.NextValue()
	}

	json, err := json.Marshal(image)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(500)
	}

	io.WriteString(w, string(json))
}

func uploadImage(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	label := r.Form.Get("label")
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + filepath.Ext(header.Filename)

	dst, err := os.Create("./public/images/" + fileName)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		log.Fatal(err)
		return
	}

	images.Insert(models.Image{Name: fileName, Label: label})
}

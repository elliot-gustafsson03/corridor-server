package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/elliot-gustafsson03/corridor-server/api/models"
)

var images = models.List{}

func AddSlideshowEndpoints() {
	http.HandleFunc("/api/get_next_image", getNextImage)
	http.HandleFunc("/api/get_all_images", getAllImages)
	http.HandleFunc("/api/upload_image", uploadImage)
	http.HandleFunc("/api/delete_image", deleteImage)

	models.LoadImages(&images)
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

func getAllImages(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, models.GenerateJson(&images))
}

func deleteImage(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	fileName := string(body)

	err = os.Remove("images/" + fileName)
	if err != nil {
		log.Fatal(err)
	}

	images.Delete(fileName)
	models.SaveImages(&images)
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

	dst, err := os.Create("images/" + fileName)
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

	newImage := models.Image{Name: fileName, Label: label}
	images.Insert(newImage)
	models.SaveImages(&images)

	io.WriteString(w, "1")
}

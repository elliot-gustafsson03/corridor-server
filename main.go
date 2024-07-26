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
	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
)

var images = models.List{}
var client *supabase.Client

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	DB_URL := os.Getenv("DB_URL")
	DB_KEY := os.Getenv("DB_KEY")

	client = createClient(DB_URL, DB_KEY)
	loadImages()
	openServer()
}

func createClient(url string, key string) *supabase.Client {
	client, err := supabase.NewClient(url, key, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func loadImages() {
	res, _, err := client.From("images").Select("*", "exact", false).Execute()
	if err != nil {
		log.Fatal(err)
	}

	var data []models.Image
	json.Unmarshal(res, &data)
	for i := 0; i < len(data); i++ {
		images.Insert(data[i])
	}
}

func openServer() {
	fs := http.FileServer(http.Dir("./public"))

	http.Handle("/", fs)
	http.HandleFunc("/api/get_next_image", getNextImage)
	http.HandleFunc("/api/get_all_images", getAllImages)
	http.HandleFunc("/api/upload_image", uploadImage)
	http.HandleFunc("/api/delete_image", deleteImage)

	port := "3333"

	log.Print("Listening on port " + port + "...")

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
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

func getAllImages(w http.ResponseWriter, r *http.Request) {
	res, _, err := client.From("images").Select("id, image, label", "exact", false).Execute()
	if err != nil {
		log.Fatal(err)
	}

	io.WriteString(w, string(res))
}

func deleteImage(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	obj := struct {
		Id int `json:"id"`
	}{}
	err = json.Unmarshal(body, &obj)
	id := obj.Id

	log.Print(id)
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

	newImage := models.Image{Name: fileName, Label: label}

	client.From("images").Insert(newImage, true, "", "minimal", "exact").Execute()
	images.Insert(newImage)

	io.WriteString(w, "1")
}

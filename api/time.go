package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func AddTimeEndpoints() {
	http.HandleFunc("/api/get_name_day", getNameDay)
}

func getNameDay(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("api/resources/namedays.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	var namedays []map[int][]string
	err = json.Unmarshal(bytes, &namedays)
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()
	month := (now.Month()) - 1
	day := now.Day()

	names := namedays[month][day]
	bytes, err = json.Marshal(names)
	if err != nil {
		log.Fatal(err)
	}

	io.WriteString(w, string(bytes))
}

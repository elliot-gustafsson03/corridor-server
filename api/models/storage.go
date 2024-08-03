package models

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
)

const FILE_PATH = "images/images.json"

func LoadImages(list *List) {
	f, err := os.Open(FILE_PATH)
	if err != nil {
		return
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	var imageArray []Image
	err = json.Unmarshal(bytes, &imageArray)

	for i := 0; i < len(imageArray); i++ {
		list.Insert(imageArray[i])
	}
}

func SaveImages(list *List) {
	writeToFile(GenerateJson(list))
}

func GenerateJson(list *List) string {
	builder := strings.Builder{}

	builder.WriteString("[\n")

	if !list.IsEmpty() {
		currentNode := list.Head

		for {
			json, err := json.Marshal(currentNode.Value)
			if err != nil {
				log.Fatal(err)
			}

			builder.Write(json)

			if currentNode.Next == list.Head {
				break
			} else {
				builder.WriteString(",\n")
				currentNode = currentNode.Next
			}
		}
	}

	builder.WriteString("\n]")

	return builder.String()
}

func writeToFile(data string) {
	f, err := os.OpenFile(FILE_PATH, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString(data)
	if err != nil {
		log.Fatal(err)
	}
}

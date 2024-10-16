package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func getRandomImagePath(dir string) (string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}

	var images []string
	for _, file := range files {
		if !file.IsDir() {
			ext := filepath.Ext(file.Name())
			if ext == ".png" || ext == ".jpg" || ext == ".jpeg" {
				images = append(images, filepath.Join(dir, file.Name()))
			}
		}
	}

	if len(images) == 0 {
		return "", fmt.Errorf("no images found in directory: %s", dir)
	}

	rand.Seed(time.Now().UnixNano())
	return images[rand.Intn(len(images))], nil
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	imagePath, err := getRandomImagePath("src/images")
	if err != nil {
		log.Println("Error getting random image:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/jpg")

	fmt.Println(imagePath)

	w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=%s", filepath.Base(imagePath)))

	http.ServeFile(w, r, imagePath)
}

func main() {
	http.HandleFunc("/", imageHandler)
	http.ListenAndServe(":8000", nil)
}

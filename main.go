package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	log.Println("Spooling up the rand() service.")
	mux := http.NewServeMux()

	// Generate a random uint64.
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		fmt.Fprintf(w, "%d", r.Uint64())
	})

	// Generate a random png.
	mux.HandleFunc("/png", func(w http.ResponseWriter, req *http.Request) {
		img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{240, 240}})
		// Each pixel consists of 4 uint8's in R,G,B,A order.
		for i := 0; i < len(img.Pix); i += 4 {
			// Generate a single uint64 rand per pixel for max speed.
			r := rand.New(rand.NewSource(time.Now().UnixNano())).Uint64()
			img.Pix[i] = uint8(r)
			img.Pix[i+1] = uint8(r >> 8)
			img.Pix[i+2] = uint8(r >> 16)
			// Always set full alpha.
			img.Pix[i+3] = 255
		}
		png.Encode(w, img)
	})

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	log.Println("Spooling up the rand() service.")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		fmt.Fprintf(w, "%d", r.Uint64())
	})
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

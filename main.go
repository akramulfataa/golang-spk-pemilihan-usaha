package main

import (
	"golang-spk-konsep/handlers"
	"log"
	"net/http"

	"fmt"
)

func main() {
	// Menggunakan FileServer untuk menangani file statis dalam folder 'public'
	fs := http.FileServer(http.Dir("public"))
	// Menggunakan StripPrefix untuk menangani root URL ('/') dan mengarahkannya ke halaman index.html
	http.Handle("/", http.StripPrefix("/", fs))
	// Handle requests to /perhitungan
    http.HandleFunc("/perhitungan", handlers.PerhitunganHandler)
	// handle handlefunc post 
	http.HandleFunc("/api/add", handlers.AddUsaha)
	// handle handlefunc get
	http.HandleFunc("/api/get", handlers.GetUsaha)
	// print port http 
	fmt.Println("Server started on port 4000")
	// listen http request
	log.Fatal(http.ListenAndServe(":4000", nil))
}

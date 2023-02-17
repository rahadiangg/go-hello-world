package main

import (
	"fmt"
	"net/http"
	"os"
)

func getAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Ini adalah halaman about")
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	http.HandleFunc("/about", getAbout)

	var port string = "3000"

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	fmt.Println("Starting web server at ", port)
	http.ListenAndServe(":"+port, nil)
}

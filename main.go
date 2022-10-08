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
		fmt.Fprintln(w, "Hello Wolrd")
	})

	http.HandleFunc("/about", getAbout)

	var port string = os.Getenv("APP_PORT")

	fmt.Println("Starting web server at ", port)
	http.ListenAndServe(":"+port, nil)
}

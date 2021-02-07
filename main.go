package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("starting server")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request recieved")
		io.WriteString(w, "Hello World")
	})

	http.HandleFunc("/end", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Killing proccess")
		io.WriteString(w, "Goodbye cruel world")
		defer os.Exit(0)
	})

	http.ListenAndServe(":80", nil)
}

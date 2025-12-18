package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})

	fmt.Println("Server is running on :8080")

	fmt.Println("hello from sajib saha")
	http.ListenAndServe(":8080", nil)
}
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

	fmt.Println("CI done or not")

	fmt.Println("hello from sajib saha -2")
	http.ListenAndServe(":8080", nil)
}
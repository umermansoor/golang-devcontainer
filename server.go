package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world from Golang devcontainer!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)

	fmt.Println("Server is running on port 3000")
}

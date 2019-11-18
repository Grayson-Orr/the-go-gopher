package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Index")
}

func setupRoutes() {
	http.HandleFunc("/", index)
}

func main() {
	fmt.Println("App listening on port 3000")
	setupRoutes()
	http.ListenAndServe(":3000", nil)
}
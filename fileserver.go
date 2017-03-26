package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	fmt.Println("File server listening on http://localhost:8123")
	err := http.ListenAndServe(":8123", nil)
	log.Fatal(err)
}

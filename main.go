package main

import (
	"GO/api/todoPkg"
	"log"
	"net/http"
)

func main() {
	router := todoPkg.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

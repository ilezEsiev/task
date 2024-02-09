package main

import (
	"log"
	"net/http"
	"task/controller"
)

func main() {
	http.HandleFunc("/users", controller.UsersHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

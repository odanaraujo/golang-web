package main

import (
	"dan/routes"
	"log"
	"net/http"
)

func main() {
	routes.LoadRoutes()
	server := http.ListenAndServe(":8000", nil)

	log.Fatal(server)
}

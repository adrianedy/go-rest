package main

import (
	"context"
	"net/http"

	"github.com/adrianedy/go-rest/database"
	"github.com/adrianedy/go-rest/router"
)

func main() {
	connection := database.GetConnection()
	defer connection.Disconnect(context.TODO())

	http.HandleFunc("/", router.Serve)

	http.ListenAndServe(":8081", nil)
}

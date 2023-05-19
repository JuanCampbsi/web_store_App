package main

import (
	"net/http"

	"store/routes"

	_ "github.com/lib/pq"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}

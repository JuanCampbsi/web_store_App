package main

import (
	"net/http"

	"projects/goStudyng/Web_Store_App/routes"

	_ "github.com/lib/pq"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}

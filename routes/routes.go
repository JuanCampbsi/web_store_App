package routes

import (
	"net/http"
	"projects/goStudyng/Web_Store_App/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
}

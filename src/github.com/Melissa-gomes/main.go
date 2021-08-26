package main

import (
	"net/http"

	"github.com/Melissa-gomes/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}

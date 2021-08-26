package routes

import (
	"net/http"

	"github.com/Melissa-gomes/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.NewProduct)
	http.HandleFunc("/insert", controllers.AddProduct)
	http.HandleFunc("/delete", controllers.DeleteProduct)
	http.HandleFunc("/edit", controllers.EditAProduct)
	http.HandleFunc("/update", controllers.Update)
}

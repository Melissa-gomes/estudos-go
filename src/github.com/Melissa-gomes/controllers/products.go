package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/Melissa-gomes/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.SearchAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}

func NewProduct(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "NewProduct", nil)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		quantity := r.FormValue("quantidade")

		priceFloat, err := strconv.ParseFloat(price, 64)

		if err != nil {
			log.Panicln("Erro na conversão do campo price:", err)
		}

		quantityInt, errQ := strconv.Atoi(quantity)

		if errQ != nil {
			log.Panicln("Erro na conversão do campo quantity:", errQ)
		}

		models.CreateNewProduct(name, description, priceFloat, quantityInt)
	}
	http.Redirect(w, r, "/", 301)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	models.ExcludeProduct(idProduct)
	http.Redirect(w, r, "/", 301)
}

func EditAProduct(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	product := models.EditedProduct(idProduct)
	temp.ExecuteTemplate(w, "Edit", product)
	//models.EditProduct(idProduct)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		quantity := r.FormValue("quantidade")

		idInt, errI := strconv.Atoi(id)

		if errI != nil {
			log.Panicln("Erro na conversão do campo Id:", errI)
		}

		priceFloat, err := strconv.ParseFloat(price, 64)

		if err != nil {
			log.Panicln("Erro na conversão do campo price:", err)
		}

		quantityInt, errQ := strconv.Atoi(quantity)

		if errQ != nil {
			log.Panicln("Erro na conversão do campo quantity:", errQ)
		}

		models.UpProduct(idInt, name, description, priceFloat, quantityInt)
	}
	http.Redirect(w, r, "/", 301)
}

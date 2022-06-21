package controllers

import (
	"dan/models"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var templatesHtml = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	allProducts := models.GetAllProducts()
	templatesHtml.ExecuteTemplate(w, "Index", allProducts)
}
func New(w http.ResponseWriter, r *http.Request) {

	templatesHtml.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		amount := r.FormValue("quantidade")

		convertPrice, err := strconv.ParseFloat(price, 64)
		isErro(err)

		convertAmount, err := strconv.Atoi(amount)
		isErro(err)

		models.SaveProduct(name, description, convertPrice, convertAmount)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idProduct)
	isErro(err)
	models.DeleteProduct(id)

	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idProduct)
	isErro(err)
	produto := models.EditProduct(id)
	templatesHtml.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		amount := r.FormValue("quantidade")

		convertPrice, err := strconv.ParseFloat(price, 64)
		isErro(err)

		convertId, err := strconv.Atoi(id)
		isErro(err)

		convertAmount, err := strconv.Atoi(amount)
		isErro(err)

		models.UpdateProduct(convertId, name, description, convertPrice, convertAmount)
		http.Redirect(w, r, "/", 301)
	}
}

func isErro(err error) {
	if err != nil {
		fmt.Println("Ocorreu um erro na conversão:", err)
		log.Fatal(err)
		panic(err.Error())
	}
}

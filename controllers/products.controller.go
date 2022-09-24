package controllers

import (
	"html/template"
	"net/http"
	"strconv"
	"webGo/main.go/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, _ *http.Request) {
	err := temp.ExecuteTemplate(w, "Index", models.GetAllProducts())
	if err != nil {
		return
	}
}

func New(w http.ResponseWriter, _ *http.Request) {
	err := temp.ExecuteTemplate(w, "New", nil)
	if err != nil {
		return
	}
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		quantity := r.FormValue("quantidade")

		newPrice, err := strconv.ParseFloat(price, 64)
		newQuantity, err := strconv.Atoi(quantity)

		if err != nil {
			panic(err.Error())
		}

		models.NewProduct(name, description, newPrice, newQuantity)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	models.DeleteProduct(productId)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	product := models.ProductEdit(productId)
	err := temp.ExecuteTemplate(w, "Edit", product)
	if err != nil {
		return
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		quantity := r.FormValue("quantidade")

		convertedId, err := strconv.Atoi(id)
		convertedPrice, err := strconv.ParseFloat(price, 64)
		convertedQuantity, err := strconv.Atoi(quantity)

		if err != nil {
			panic(err.Error())
		}

		models.ProductUpdate(convertedId, name, description, convertedPrice, convertedQuantity)
	}

	http.Redirect(w, r, "/", 301)
}

package models

import (
	"database/sql"
	"webGo/main.go/db"
)

type Product struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func GetAllProducts() []Product {
	database := db.DatabaseConnect()

	output, err := database.Query("select * from products order by id asc")
	if err != nil {
		panic(err.Error())
	}

	typeInstance := Product{}
	var products []Product

	for output.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = output.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}

		typeInstance.Id = id
		typeInstance.Preco = price
		typeInstance.Nome = name
		typeInstance.Quantidade = quantity
		typeInstance.Descricao = description

		products = append(products, typeInstance)

	}

	defer func(database *sql.DB) {
		err := database.Close()
		if err != nil {

		}
	}(database)
	return products
}

func NewProduct(name, description string, price float64, quantity int) {
	database := db.DatabaseConnect()

	insert, err := database.Prepare("insert into products(name, description, price, quantity) values ($1, $2, $3 ,$4)")

	if err != nil {
		panic(err.Error())
	}

	_, err = insert.Exec(name, description, price, quantity)
	if err != nil {
		return
	}

	defer func(database *sql.DB) {
		err := database.Close()
		if err != nil {

		}
	}(database)
}

func DeleteProduct(id string) {
	database := db.DatabaseConnect()
	deleteProduct, err := database.Prepare("delete from products where id = $1")

	if err != nil {
		panic(err.Error())
	}

	_, err = deleteProduct.Exec(id)
	if err != nil {
		return
	}

	defer func(database *sql.DB) {
		err := database.Close()
		if err != nil {

		}
	}(database)
}

func ProductEdit(id string) Product {
	database := db.DatabaseConnect()

	output, err := database.Query("select * from products where id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	productToUpdate := Product{}

	for output.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = output.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}

		productToUpdate.Id = id
		productToUpdate.Nome = name
		productToUpdate.Quantidade = quantity
		productToUpdate.Preco = price
		productToUpdate.Descricao = description
	}

	defer func(database *sql.DB) {
		err := database.Close()
		if err != nil {

		}
	}(database)
	return productToUpdate
}

func ProductUpdate(id int, name, description string, price float64, quantity int) {
	database := db.DatabaseConnect()

	output, err := database.Prepare("update products set name=$1, description=$2, price=$3, quantity=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}

	_, err = output.Exec(name, description, price, quantity, id)
	if err != nil {
		return
	}
	defer func(database *sql.DB) {
		err := database.Close()
		if err != nil {

		}
	}(database)
}

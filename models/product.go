package models

import (
	"dan/bd"
	"fmt"
	"log"
)

type Product struct {
	Name, Description string
	Price             float64
	Amount, Id        int
}

func GetAllProducts() []Product {
	db := bd.ConnectionDB()
	selectAllProducts, err := db.Query("SELECT * FROM products order by id asc")

	product := Product{}
	products := []Product{}
	isErro(err)

	for selectAllProducts.Next() {
		var id int
		var name, description string
		var price float64
		var amount int

		err = selectAllProducts.Scan(&id, &name, &description, &price, &amount)
		isErro(err)
		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Amount = amount

		products = append(products, product)
	}
	defer db.Close()
	return products
}

func SaveProduct(name, description string, price float64, amount int) {

	bd := bd.ConnectionDB()

	saveProduct, err := bd.Prepare("INSERT INTO products(name, description, price, amount)" +
		"values($1, $2, $3, $4)")
	isErro(err)
	saveProduct.Exec(name, description, price, amount)
	defer bd.Close()
}

func DeleteProduct(id int) {
	bd := bd.ConnectionDB()

	deleteProduct, err := bd.Prepare("delete from products where id = $1")
	isErro(err)
	deleteProduct.Exec(id)

	defer bd.Close()
}

func EditProduct(id int) Product {
	bd := bd.ConnectionDB()

	editProduct, err := bd.Query("SELECT * FROM products where id = $1", id)
	isErro(err)
	product := Product{}
	for editProduct.Next() {
		var name, description string
		var price float64
		var amount int
		err = editProduct.Scan(&id, &name, &description, &price, &amount)
		isErro(err)
		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Amount = amount
	}
	defer bd.Close()
	return product
}

func UpdateProduct(id int, name, description string, price float64, amount int) {

	bd := bd.ConnectionDB()

	updateProduct, err := bd.Prepare("update products set name = $1, description=$2, price=$3, amount=$4 where id=$5")
	isErro(err)
	updateProduct.Exec(name, description, price, amount, id)
	defer bd.Close()
}

func isErro(err error) {
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		log.Fatal(err)
	}
}

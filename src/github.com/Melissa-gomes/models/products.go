package models

import (
	"github.com/Melissa-gomes/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func SearchAllProducts() []Product {
	db := db.ConnectWithDataBase()

	selectAllProducts, err := db.Query("select * from products order by id asc")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectAllProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}

	defer db.Close()

	return products
}

func CreateNewProduct(name, description string, priceFloat float64, quantityInt int) {
	db := db.ConnectWithDataBase()

	insertDataInDataBase, err := db.Prepare("insert into products (name, description, price, quantity) values ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insertDataInDataBase.Exec(name, description, priceFloat, quantityInt)

	defer db.Close()
}

func ExcludeProduct(id string) {
	db := db.ConnectWithDataBase()

	delete, err := db.Prepare("delete from products where id=$1")

	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)

	defer db.Close()
}

func EditedProduct(id string) Product {
	db := db.ConnectWithDataBase()

	productFromTheDataBase, err := db.Query("select * from products where id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	productForUpdate := Product{}

	for productFromTheDataBase.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productFromTheDataBase.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}

		productForUpdate.Id = id
		productForUpdate.Name = name
		productForUpdate.Description = description
		productForUpdate.Price = price
		productForUpdate.Quantity = quantity

	}

	defer db.Close()

	return productForUpdate
}

func UpProduct(idInt int, name, description string, priceFloat float64, quantityInt int) {
	db := db.ConnectWithDataBase()

	attProduct, err := db.Prepare("update products set name=$1, description=$2, price=$3, quantity=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}

	attProduct.Exec(name, description, priceFloat, quantityInt, idInt)

	defer db.Close()
}

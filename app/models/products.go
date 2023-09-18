package models

import "go-store/database"

type Product struct {
	Id, Amount        int
	Name, Description string
	Price             float64
}

func FindAllProducts() []Product {
	database := db.DatabaseConnection()

	selectAllProducts, error := database.Query("select * from products order by id asc")
	if error != nil {
		panic(error.Error())
	}

	product := Product{}
	products := []Product{}

	for selectAllProducts.Next() {
		var id, amount int
		var name, description string
		var price float64

		error = selectAllProducts.Scan(&id, &name, &price, &description, &amount) // This data must follow the db table order
		if error != nil {
			panic(error.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Amount = amount

		products = append(products, product)
	}
	defer database.Close()
	return products
}

func CreateProduct(name, description string, price float64, amount int) {
  database := db.DatabaseConnection()
	
	newProduct, error := database.Prepare("insert into products(name, description, price, amount) values($1, $2, $3, $4)")
	if error != nil {
		panic(error.Error())
	}

	newProduct.Exec(name, description, price, amount)
	defer database.Close()
}

func DeleteProduct(id string) {
	database := db.DatabaseConnection()

	deleteProduct, error := database.Prepare("delete from products where id=$1")
	if error != nil {
		panic(error.Error())
	}

	deleteProduct.Exec(id)
	defer database.Close()
}

func EditProduct(id string) Product{
	database := db.DatabaseConnection()

	product, error := db.DatabaseConnection().Query("select * from products where id=$1", id)

	if error != nil {
		panic(error.Error())
	}

	productToUpdate := Product{}

	for product.Next() {
		var id, amount int
		var name, description string
		var price float64

		error = product.Scan(&id, &name, &price, &description, &amount)
		if error != nil {
			panic(error.Error())
		}

		productToUpdate.Id = id
		productToUpdate.Name = name
		productToUpdate.Description = description
		productToUpdate.Amount = amount
		productToUpdate.Price = price
	}

	defer database.Close()
	return productToUpdate
}

func UpdateProduct(id int, name string, price float64, description string, amount int) {
	database := db.DatabaseConnection()

	updateProduct, error := database.Prepare("update products set name=$1, description=$2, price=$3, amount=$4 where id=$5")

	if error != nil {
		panic(error.Error())
	}

	updateProduct.Exec(name, description, price, amount, id)
	defer database.Close()
}
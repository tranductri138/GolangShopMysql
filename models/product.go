package models

import (
	"encoding/json"
	"log"
	"net/http"
	"tris/db"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	IsSales     bool    `json:"is_sale"`
	Price       float32 `json:"price"`
	PriceSale   float32 `json:"price_sale"`
	Description string  `json:"des"`
	CreatedAt   int64   `json:"created_at"`
	ModifiedAt  int64   `json:"modified_at"`
	CategoryId  int     `json:"category_id"`
}

func indexByID(products []Product, id int) int {
	for i := 0; i < len(products); i++ {
		if products[i].Id == id {
			return i
		}
	}
	return -1
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products []Product
	result, err := db.Db.Query("SELECT * from products")
	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()

	for result.Next() {
		var product Product
		err := result.Scan(&product.Id, &product.Name, &product.IsSales,
			&product.Description, &product.Price, &product.PriceSale)
		if err != nil {
			panic(err.Error())
		}
		products = append(products, product)
	}
	json.NewEncoder(w).Encode(products)
}

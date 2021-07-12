package routers

import (
	products "tris/models"

	"github.com/gorilla/mux"
)

func ConnectRouter(router *mux.Router) {
	// product
	router.HandleFunc("/products", products.GetProducts).Methods("GET")
	// router.HandleFunc("/product", createPost).Methods("POST")
	// router.HandleFunc("/product/{id}", getPost).Methods("GET")
	// router.HandleFunc("/product/{id}", updatePost).Methods("PUT")
	// router.HandleFunc("/product/{id}", deletePost).Methods("DELETE")
}

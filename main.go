package main

import (
	"net/http"
	"tris/db"
	"tris/routers"

	"github.com/gorilla/mux"
)

func main() {
	db.Connect()
	r := mux.NewRouter()
	routers.ConnectRouter(r)

	http.ListenAndServe(":3000", r)
}

package main

import (
	"github.com/KaiKaizxc/Golang_pro/Golang_book_store/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9000", r))

}

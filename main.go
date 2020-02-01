package main

import (
	"OrderService/handler"

	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

var (
	tx *sqlx.Tx
	db *sqlx.DB
)

func main() {
	log.Println("db Init Strat")
	_, errMsg, err := handler.DbInit()
	if err != nil {
		log.Println(errMsg)
		log.Fatal(err)
	}
	log.Println("db Init Complete")
	log.Println("Server Start...")
	r := mux.NewRouter()
	r.HandleFunc("/", handler.MenuPage).Methods("GET")
	r.HandleFunc("/AddMenu", handler.AddMenu).Methods("POST")
	r.HandleFunc("/EditMenu", handler.EditMenu).Methods("PUT")
	r.HandleFunc("/DeleteMenu", handler.DeleteMenu).Methods("DELETE")
	http.ListenAndServe(":3000", r)

}

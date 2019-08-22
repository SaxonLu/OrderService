package main

import (
	"encoding/json"
	"net/http"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/gorilla/mux"
)

type Menu struct {
	Menu_ID int `db:"Menu_ID"`
	Menu_Name string `db:"Menu_Name"`
	Price int `db:"Price"`
	OnWork  string `db:"OnWork"`
}

func MenuPage(rw http.ResponseWriter, request *http.Request) {

	conn, err:= sqlx.Connect("mysql","sa:qwer!1234@tcp(127.0.0.1:3306)/orderservice")
	if err != nil{
		log.Fatalln(err)
	}
	defer conn.Close()
	Menu:=[]Menu{}
	conn.Select(&Menu,"select * from menu")

	b, err := json.Marshal(Menu)

	if err != nil {
		return
	}
	rw.Header().Set("Content-Type", "application/json;charset=UTF-8")
	rw.WriteHeader(http.StatusOK)
	rw.Write(b)
}

func AddMenu(rw http.ResponseWriter, request *http.Request) {
	// rw.Write([]byte("Welcome ! type :3000/Menu to seen Menu"))

	// data := make(map[string]interface{})
	data :=&Menu{}
	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&data)

	if err != nil {
		rw.Write([]byte("Error! Invalid request payload"))
		return
	}
	defer request.Body.Close()

	b, err := json.Marshal(data)

	if err != nil {
		return
	}
	rw.Header().Set("Content-Type", "application/json;charset=UTF-8")
	rw.WriteHeader(http.StatusOK)
	rw.Write(b)

}


func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", MenuPage).Methods("GET")
	r.HandleFunc("/AddMenu",AddMenu).Methods("POST")
	http.ListenAndServe(":3000", r)
	
}

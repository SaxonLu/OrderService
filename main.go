package main

import (
	// "encoding/json"
	// "net/http"
	// "github.com/gorilla/mux"

	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// import (
// 	_ "github.com/mattn/go-adodb"
// )

type Menu struct {
	Menu_ID int `db:"Menu_ID"`
	Menu_Name string `db:"Menu_Name"`
	Price int `db:"Price"`
	OnWork  string `db:"OnWork"`
}

// func MenuPage(rw http.ResponseWriter, request *http.Request) {

// 	u := &Menu{
// 		Meal:  "Steak",
// 		Price: 199,
// 	}
// 	b, err := json.Marshal(u)

// 	if err != nil {
// 		return
// 	}
// 	rw.Header().Set("Content-Type", "application/json;charset=UTF-8")
// 	rw.WriteHeader(http.StatusOK)
// 	rw.Write(b)
// }

// func Welcome(rw http.ResponseWriter, request *http.Request) {
// 	rw.Write([]byte("Welcome ! type :3000/Menu to seen Menu"))
// }


func main() {

	// r := mux.NewRouter()
	// r.HandleFunc("/", Welcome).Methods("GET")
	// r.HandleFunc("/Menu", MenuPage).Methods("GET")

	// http.ListenAndServe(":3000", r)
	conn, err:= sqlx.Connect("mysql","sa:qwer!1234@tcp(127.0.0.1:3306)/orderservice")
	if err != nil{
		log.Fatalln(err)
	}

	Menu:=[]Menu{}
	conn.Select(&Menu,"select * from menu")
	log.Println("Menu...")
	log.Println(Menu)
}

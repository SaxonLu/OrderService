package main

import (
	"OrderService/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

var tx *sqlx.Tx

func MenuPage(rw http.ResponseWriter, request *http.Request) {

	conn, err := sqlx.Connect("mysql", "sa:qwer!1234@tcp(127.0.0.1:3306)/orderservice")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	Menu := []model.Menu{}
	conn.Select(&Menu, "select * from menu")

	b, err := json.Marshal(Menu)

	if err != nil {
		return
	}
	rw.Header().Set("Content-Type", "application/json;charset=UTF-8")
	rw.WriteHeader(http.StatusOK)
	rw.Write(b)
}

func AddMenu(rw http.ResponseWriter, request *http.Request) {

	data := &model.Menu{}

	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&data)

	if err != nil {
		rw.Write([]byte("錯誤! 無法辨識的請求!"))
		return
	}

	pk := GetMenuNewPK()

	conn, err := sqlx.Connect("mysql", "sa:qwer!1234@tcp(127.0.0.1:3306)/orderservice")
	if err != nil {
		rw.Write([]byte("錯誤! mysql連線失敗!"))
	}

	tx, err = conn.Beginx()

	if err != nil {
		rw.Write([]byte("錯誤! mysql交易啟動失敗!"))
	}
	_, err = tx.Exec("INSERT INTO menu(Menu_ID, Menu_Name, Price, OnWork)values(?,?,?,?)", pk, data.Menu_Name, data.Price, data.OnWork)

	if err != nil {
		tx.Rollback()
		rw.Write([]byte("錯誤! mysql交易執行失敗!"))
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		rw.Write([]byte("錯誤! mysql交易執行失敗!"))
	}

	rw.Header().Set("Content-Type", "application/json;charset=UTF-8")

	rw.WriteHeader(http.StatusOK)

	rw.Write([]byte("菜單新增完成!"))

}

func EditMenu(rw http.ResponseWriter, request *http.Request) {
	data := &model.Menu{}

	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&data)

	if err != nil {
		rw.Write([]byte("錯誤! 無法辨識的請求!"))
		return
	}

	check := GetSingleMenu(data.Menu_ID)

	if check == false {
		rw.Write([]byte("錯誤! 資料庫內無對應菜單!"))
		return
	}

	conn, err := sqlx.Connect("mysql", "sa:qwer!1234@tcp(127.0.0.1:3306)/orderservice")
	if err != nil {
		rw.Write([]byte("錯誤! mysql連線失敗!"))
		return
	}

	tx, err = conn.Beginx()

	if err != nil {
		rw.Write([]byte("錯誤! mysql交易啟動失敗!"))
		return
	}

	_, err = tx.Exec("UPDATE menu SET Menu_Name=?, Price=?, OnWork=? WHERE Menu_ID=?", data.Menu_Name, data.Price, data.OnWork, data.Menu_ID)

	if err != nil {
		tx.Rollback()
		rw.Write([]byte("錯誤! mysql交易執行失敗!"))
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		rw.Write([]byte("錯誤! mysql交易執行失敗!"))
	}

	rw.Header().Set("Content-Type", "application/json;charset=UTF-8")

	rw.WriteHeader(http.StatusOK)

	rw.Write([]byte("菜單更新完成!"))

}

func DeleteMenu(rw http.ResponseWriter, request *http.Request) {
	data := &model.Menu{}

	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&data)

	if err != nil {
		rw.Write([]byte("錯誤! 無法辨識的請求!"))
		return
	}

	check := GetSingleMenu(data.Menu_ID)

	if check == false {
		rw.Write([]byte("錯誤! 資料庫內無對應菜單!"))
		return
	}

	conn, err := sqlx.Connect("mysql", "sa:qwer!1234@tcp(127.0.0.1:3306)/orderservice")
	if err != nil {
		rw.Write([]byte("錯誤! mysql連線失敗!"))
		return
	}

	tx, err = conn.Beginx()

	if err != nil {
		rw.Write([]byte("錯誤! mysql交易啟動失敗!"))
		return
	}

	_, err = tx.Exec("DELETE FROM menu WHERE Menu_ID=?", data.Menu_ID)

	if err != nil {
		tx.Rollback()
		rw.Write([]byte("錯誤! mysql交易執行失敗!"))
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		rw.Write([]byte("錯誤! mysql交易執行失敗!"))
	}

	rw.Header().Set("Content-Type", "application/json;charset=UTF-8")

	rw.WriteHeader(http.StatusOK)

	rw.Write([]byte("菜單刪除完成!"))
}

func GetMenuNewPK() int {
	conn, err := sqlx.Connect("mysql", "sa:qwer!1234@tcp(127.0.0.1:3306)/orderservice")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	result := 0
	err = conn.Get(&result, "select (MAX(Menu_ID)+1) from menu")
	if err != nil {
	}
	return result
}

func GetSingleMenu(pk int) bool {
	conn, err := sqlx.Connect("mysql", "sa:qwer!1234@tcp(127.0.0.1:3306)/orderservice")
	if err != nil {
		return false
	}

	defer conn.Close()

	result := model.Menu{}

	err = conn.Get(&result, "SELECT * FROM menu where Menu_ID=? ", pk)

	if err != nil {

		fmt.Println(err)

		return false
	}

	return true
}

func main() {

	log.Println("db Init Strat")
	if err := dbInit(); err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/", MenuPage).Methods("GET")
	r.HandleFunc("/AddMenu", AddMenu).Methods("POST")
	r.HandleFunc("/EditMenu", EditMenu).Methods("PUT")
	r.HandleFunc("/DeleteMenu", DeleteMenu).Methods("DELETE")
	http.ListenAndServe(":3000", r)

}

func dbInit() error {

	return nil
}

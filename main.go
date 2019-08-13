package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Menu struct {
	Meal  string
	Price int
}

func MenuPage(rw http.ResponseWriter, request *http.Request) {

	u := &Menu{
		Meal:  "Steak",
		Price: 199,
	}
	b, err := json.Marshal(u)

	if err != nil {
		return
	}
	rw.Header().Set("Content-Type", "application/json;charset=UTF-8")
	rw.WriteHeader(http.StatusOK)
	rw.Write(b)
}

func Welcome(rw http.ResponseWriter, request *http.Request) {
	rw.Write([]byte("Welcome ! type :3000/Menu to seen Menu"))
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", Welcome).Methods("GET")
	r.HandleFunc("/Menu", MenuPage).Methods("GET")

	http.ListenAndServe(":3000", r)
}

package handler

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func DbInit() (db *sqlx.DB, status string, err error) {
	db, err = sqlx.Open("mysql", "sa:qwer!1234@tcp(127.0.0.1:3306)/orderservice")
	if err != nil {
		status = "錯誤! mysql連線失敗!"
		return
	}
	if err = db.Ping(); err != nil {
		status = "錯誤! mysql連線失敗!"
		return
	}
	return
}

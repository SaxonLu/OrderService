package model

type Menu struct {
	Menu_ID   int    `db:"Menu_ID"`
	Menu_Name string `db:"Menu_Name"`
	Price     int    `db:"Price"`
	OnWork    string `db:"OnWork"`
}

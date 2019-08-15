package main

import (
	// "encoding/json"
	// "net/http"
	// "github.com/gorilla/mux"

	"database/sql"
	"flag"
	"fmt"
	"log"
)

import (
	_ "github.com/mattn/go-adodb"
)

// type Menu struct {
// 	Meal  string
// 	Price int
// }

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


var (
	local    bool
	remoteIP string
	remoteDS string
)

func init() {
	flag.BoolVar(&local, "local", true, "set window connect.")
	flag.StringVar(&remoteIP, "remoteIP", "192.168.1.104", "set up remote mssql of ip.")
	flag.StringVar(&remoteDS, "remoteDS", "MSSQLSERVER", "set up remote mssql of datasource.")
}

type Mssql struct {
	*sql.DB
	dataSource string
	database   string
	windows    bool
	sa         *SA
}

type SA struct {
	user   string
	passwd string
	port   int
}

func NewMssql() *Mssql {
	mssql := new(Mssql)
	dataS := "(localdb)\v11.0\\MSSQLSERVER"
	if !local {
		dataS = fmt.Sprintf("%s\\%s", remoteIP, remoteDS)
	}

	mssql = &Mssql{
		// 如果数据库是默认实例（MSSQLSERVER）则直接使用IP，命名实例需要指明。
		// dataSource: "192.168.1.104\\MSSQLSERVER",
		dataSource: dataS,
		database:   "Northwind",
	    windows: false ,//为windows身份验证，false 必须设置sa账号和密码
		//windows: local,
		sa: &SA{
			user:   "sa",
			passwd: "qwer!1234",
			port:   1433,
		},
	}

	fmt.Println(mssql)

	return mssql

}

func (m *Mssql) Open() error {
	config := fmt.Sprintf("Provider=SQLOLEDB;Initial Catalog=%s;Data Source=%s",
		m.database, m.dataSource)

	if m.windows {
		config = fmt.Sprintf("%s;Integrated Security=SSPI", config)
	}else {
		// sql 2000的端口写法和sql 2005以上的有所不同，在Data Source 后以逗号隔开。
		config = fmt.Sprintf("%s;user id=%s;password=%s",
			config, m.sa.user, m.sa.passwd)
	}


	var err error
	m.DB, err = sql.Open("adodb", config)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(config)

	return err
}

func (m *Mssql) Select() {
	rows, err := m.Query("select GETDATE()")
	if err != nil {
		fmt.Printf("select query err: %s\n", err)
	}

	for rows.Next() {
		var id, name string
		rows.Scan(&id, &name)
		fmt.Printf("LastName = %s, FirstName = %s\n", id, name)
	}
}

func main() {

	// r := mux.NewRouter()
	// r.HandleFunc("/", Welcome).Methods("GET")
	// r.HandleFunc("/Menu", MenuPage).Methods("GET")

	// http.ListenAndServe(":3000", r)

	flag.Parse()

	mssql := NewMssql()
	err := mssql.Open()
	checkError(err)

	mssql.Select()
}
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
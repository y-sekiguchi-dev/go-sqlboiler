package sqlboiler

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/boil"
)

type Closable func()

type mysqlInfo struct {
	host   string
	port   uint
	dbname string
	usr    string
	pass   string
}

func (m mysqlInfo) connectionString() string {
	login := fmt.Sprintf("%s:%s", m.usr, m.pass)
	db := fmt.Sprintf("tcp(%s:%d)/%s", m.host, m.port, m.dbname)
	return login + "@" + db
}

func Init() Closable {
	dbInfo := mysqlInfo{
		host:   "localhost",
		port:   3306,
		dbname: "test_database",
		usr:    "docker",
		pass:   "docker",
	}.connectionString()
	println(dbInfo)
	db, err := sql.Open("mysql", dbInfo)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	boil.SetDB(db)
	return func() {
		db.Close()
	}
}

package driver

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// DB ...
type DB struct {
	SQL *sql.DB
	// Mgo *mgo.database
}

// DBConn ...
var dbConn = &DB{}

// ConnectSQL ...
func ConnectSQL(host, port, uname, pass, dbname string) (*DB, error) {
	/*dbSource := fmt.Sprintf(
		"root:%s@tcp(%s:%s)/%s?charset=utf8",
		pass,
		host,
		port,
		dbname,
	)
	*/
	d, err := sql.Open("mysql", "root:123@tcp(10.43.238.33:3005)/userManagement")
	if err != nil {
		panic(err)
	}
	dbConn.SQL = d
	return dbConn, err
}

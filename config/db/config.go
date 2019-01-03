package db

import(
	"demo/config"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"fmt"
)

var Conn *sql.DB

var err error

type dbConf struct{
	database string
	host string
	username string
	password string
	port int
}

var DbConf dbConf

func init(){
	initDbConf()
	fmt.Println("Initialization the database connection")
	Conn, err = sql.Open("mysql", buildconnectionString())	
	config.RaiseException(err)
	err = Conn.Ping()
	config.RaiseException(err)
}


func ExecSql(sqlString string) *sql.Rows{
	rows, err := Conn.Query(sqlString)
	config.RaiseException(err)
	return rows
}

func initDbConf(){
	var c config.Conf
	data := c.FetchConfig("/Users/unbxd/go-workspace/src/demo/config/db/config.yml")
	DbConf = dbConf{
			data["database"].(string),
			data["host"].(string),
			data["username"].(string),
			data["password"].(string),
			data["port"].(int),
		}
}

func buildconnectionString() string{
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", DbConf.username, DbConf.password, DbConf.host, DbConf.port, DbConf.database)
}
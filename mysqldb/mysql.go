package mysqldb

import (
	"flag"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Pool   *sqlx.DB
	Schema string
)

var (
	User     string
	Host     string
	Password string
)

func init() {
	var usage = "target=xxoo go run main.go -u user -h 127.0.0.1 -p xxoo -d database"
	user := flag.String("u", "", "sql user")
	host := flag.String("h", "127.0.0.1", "sql host")
	password := flag.String("p", "", "sql password")
	database := flag.String("d", "", "sql database")
	flag.Parse()
	if *user == "" {
		log.Fatalf("user is nil. usage is: %s", usage)
	}
	if *password == "" {
		log.Fatalf("password is nil. usage is: %s", usage)
	}
	if *database == "" {
		log.Fatalf("database is nil. usage is: %s", usage)
	}
	User = *user
	Host = *host
	Password = *password
	Schema = *database
	var cc = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", User, Password, Host, Schema)
	Pool = sqlx.MustOpen("mysql", cc)
	Pool.SetMaxOpenConns(1000)
	Pool.SetMaxIdleConns(10)
}

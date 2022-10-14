package mysql_connection

import (
	"database/sql"
	"fmt"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// dont forgot to run: go get github.com/go-sql-driver/mysql
func InitMysqlDB() *sql.DB {
	var (
		errMysql error
		dbConn   *sql.DB
	)

	dbHost := "localhost"
	dbPort := "3306"
	dbUser := "root"
	dbPass := ""
	dbName := "salt_academy_exam_2"

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	dbConn, errMysql = sql.Open(`mysql`, dsn)

	if errMysql != nil {
		panic(errMysql)
	}

	dbConn.SetMaxOpenConns(300)
	dbConn.SetMaxIdleConns(25)
	dbConn.SetConnMaxLifetime(5 * time.Minute)

	return dbConn
}

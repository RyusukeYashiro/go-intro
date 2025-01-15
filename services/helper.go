package services

import (
	"database/sql"
	"fmt"
	"os"
)

var (
	dbUser = os.Getenv("USER_NAME")      // DB_USER から USER_NAME に変更
    dbPassword = os.Getenv("USER_PASS")  // DB_PASSWORD から USER_PASS に変更
    dbDatabase = os.Getenv("DATABASE")   // DB_NAME から DATABASE に変更
	dbConn = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)	
)

func connectDB() (*sql.DB , error) {
	db , err := sql.Open("mysql" , dbConn)
	if err != nil {
		return nil , err
	}
	return db , nil
}
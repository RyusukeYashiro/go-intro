package repositories_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
)

var testDB *sql.DB

func setup() error {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser,dbPassword, dbDatabase)
	var err error
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
	return err
	}

	return nil
}

func teardown() {
	testDB.Close()
}

// TestMain 関数を定義することによって、「前処理→ユニットテスト→後処理」というフローを作る
func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		os.Exit(1)
	}
	
	m.Run()

	teardown()
}
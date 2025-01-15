package repositories_test

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"testing"
)

var testDB *sql.DB

var (
	dbUser = os.Getenv("USER_NAME")      // DB_USER から USER_NAME に変更
    dbPassword = os.Getenv("USER_PASS")  // DB_PASSWORD から USER_PASS に変更
    dbDatabase = os.Getenv("DATABASE")   // DB_NAME から DATABASE に変更
	dbConn = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser,dbPassword, dbDatabase)
)

func connectDB () error {
	var err error
	testDB , err = sql.Open("mysql" , dbConn)
	if  err != nil {
		fmt.Printf("this is connectError");
		return err
	}
	return nil
}

func setup() error {
	if err := connectDB(); err != nil {
		fmt.Println(err)
		return err
	}	
	if err := CleanupDB(); err != nil {
		fmt.Println("cleanup")
		return err
	}
	if err := setupTestData(); err != nil {
		fmt.Println("setup")
		return err
	}
	return nil
}

func teardown() {
	CleanupDB()
	testDB.Close()
}

func setupTestData() error {
	//mysqlのクエリを実行させるexecパッケージのCommandを採用
	cmd := exec.Command("mysql" , "-h" , "127.0.0.1" , "-u" , "docker" , "sampledb" , 
	"--password=docker" , "-e" , "source ./testdata/setupDB.sql")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Command execution failed: %s\nOutput: %s\n", err, string(output))
		return err
	}
	err2 := cmd.Run()
	if err2 != nil {
		return err
	}
	return nil
}

func CleanupDB() error {
	cmd := exec.Command("mysql", "-h", "127.0.0.1", "-u", "docker", "sampledb",
	"--password=docker" , "-e", "source ./testdata/cleanupDB.sql")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
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
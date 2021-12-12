package config

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"

	_ "github.com/go-sql-driver/mysql"
)

func initSQLDocker() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", os.Getenv("CONN_STRING"))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitSQL() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "KlikA2C:KlikA2C123!@tcp(127.0.0.1:3306)/KlikA2C?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db, err := initSQLDocker()
		return db, err
		// return nil, err
	}
	return db, nil
}

func InitSQLTest() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "KlikA2C:KlikA2C123!@tcp(127.0.0.1:3306)/KlikA2C_Test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Migrate(name string) {
	db, err := InitSQL()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE if not exists " + name)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("/bin/sh", "init.sh")
	err = cmd.Run()

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("errrror")
		return
	}
}

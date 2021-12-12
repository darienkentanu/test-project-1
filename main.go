package main

import (
	"os"
	"test-project-1/config"
	"test-project-1/route"
)

func main() {
	if os.Getenv("CONN_STRING") == "" {
		config.Migrate("KlikA2C;")
		db, err := config.InitSQL()
		if err != nil {
			panic(err)
		}
		e := route.NewRoutes(db)
		e.Logger.Fatal(e.Start(":8080"))
	}
	docker()
}

func docker() {
	config.Migrate("KlikA2C;")
	db, err := config.InitSQLDocker()
	if err != nil {
		panic(err)
	}
	e := route.NewRoutes(db)
	e.Logger.Fatal(e.Start(":8080"))
}

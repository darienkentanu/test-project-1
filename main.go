package main

import (
	"test-project-1/config"
	"test-project-1/helper"
	"test-project-1/route"
)

func main() {
	config.Migrate("KlikA2C")
	db, err := helper.InitDB()
	if err != nil {
		panic(err)
	}
	e := route.NewRoutes(db)
	e.Logger.Fatal(e.Start(":8080"))
}

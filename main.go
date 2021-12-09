package main

import (
	"test-project-1/config"
	"test-project-1/route"
)

func main() {
	db, err := config.InitSQL()
	if err != nil {
		panic(err)
	}
	e := route.NewRoutes(db)
	e.Logger.Fatal(e.Start(":8080"))
}

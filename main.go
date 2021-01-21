package main

import (
	"golang-crud-test/routes"
	"golang-crud-test/db"
)

func main() {
	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":9001"))
}

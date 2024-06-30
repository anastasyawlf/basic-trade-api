package main

import (
	"basic-trade/database"
	"basic-trade/router"
)

var PORT = ":8080"

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(PORT)
}

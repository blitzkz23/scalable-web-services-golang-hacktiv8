package main

import (
	"scalable-web-services-golang-hacktiv8/module_10/JWT/database"
	"scalable-web-services-golang-hacktiv8/module_10/JWT/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}

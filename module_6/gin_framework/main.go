package main

import "scalable-web-services-golang-hacktiv8/module_6/gin_framework/routers"

// !  go get -u github.com/gin-gonic/gin
func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}

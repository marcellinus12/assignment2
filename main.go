package main

import (
	"fmt"
	"web-server/database"
	"web-server/routers"
)

func main() {
	database.StartDB()
	const PORT = ":8088"

	fmt.Sprintln("server start at", PORT)
	routers.StartServer().Run(PORT)
}

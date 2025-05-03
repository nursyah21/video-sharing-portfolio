package main

import "video-go/database"

func main() {
	database.ConnectDB()
	database.Drop()
	database.Migrate()
	defer database.CloseDB()
}

package main

import "video-go/database"

func main() {
	database.ConnectDB()
	database.Migrate()
	defer database.CloseDB()
}

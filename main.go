package main

import (
	db "server/system/db"
	server "server/system/server"
	// "server/controllers"

)

func main() {
	db.ConnectToDB()
	server.LaunchServer()
	// controllers.DeleteAll()
	db.StopDb()
}

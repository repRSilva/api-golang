package main

import (
	"github.com/repRSilva/api-golang/database"
	"github.com/repRSilva/api-golang/server"
)

func main() {

	database.StartDB()
	server := server.NewServer()
	server.Run()
}

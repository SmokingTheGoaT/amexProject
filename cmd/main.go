package main

import (
	"amexProject/app"
	"amexProject/handler"
	"amexProject/repository"
)

func main(){
	database := repository.InitialiseDB()
	services := app.New(database)
	server := handler.New(services)
	server.InitialiseRoutes()
	server.Run()
}
package main

import (
	"log"

	"github.com/GicGa-iOS/prueba-twitter/bd"
	"github.com/GicGa-iOS/prueba-twitter/handlers"
)

func main() {
	if !bd.CheckConnection() {
		log.Fatal("----No connection to the database----")
		return
	}
	handlers.Handlers()
}

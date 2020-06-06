package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/GicGa-iOS/prueba-twitter/middlew"
	"github.com/GicGa-iOS/prueba-twitter/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Handlers setting port, handler and start listenning to server*/
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8081"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}

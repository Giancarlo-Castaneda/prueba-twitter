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
	router.HandleFunc("/showProfile", middlew.CheckDB(middlew.ValidJWT(routers.ShowProfile))).Methods("GET")
	router.HandleFunc("/updateProfile", middlew.CheckDB(middlew.ValidJWT(routers.UpdateProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.CheckDB(middlew.ValidJWT(routers.SaveTweet))).Methods("POST")
	router.HandleFunc("/readTweets", middlew.CheckDB(middlew.ValidJWT(routers.ReadTweets))).Methods("GET")
	router.HandleFunc("/deleteTweet", middlew.CheckDB(middlew.ValidJWT(routers.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/uploadAvatar", middlew.CheckDB(middlew.ValidJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/getAvatar", middlew.CheckDB(routers.GetAvatar)).Methods("GET")
	router.HandleFunc("/uploadBanner", middlew.CheckDB(middlew.ValidJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/getBanner", middlew.CheckDB(routers.GetBanner)).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}

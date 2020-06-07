package routers

import (
	"encoding/json"
	"net/http"

	"github.com/GicGa-iOS/prueba-twitter/bd"
	"github.com/GicGa-iOS/prueba-twitter/models"
)

/*Register is the function to create the user registry in the database*/
func Register(w http.ResponseWriter, r *http.Request) {

	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error in the received data: "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "User email is required", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Please specify a password of at least 6 characters", 400)
		return
	}

	_, found, _ := bd.CheckUserExists(t.Email)
	if found == true {
		http.Error(w, "There is already a registered user with that email", 400)
		return
	}

	_, status, err := bd.InsertRegister(t)
	if err != nil {
		http.Error(w, "An error occurred while trying to register the user: "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Could not insert user registry", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

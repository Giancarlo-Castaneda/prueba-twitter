package routers

import (
	"encoding/json"
	"net/http"

	"github.com/GicGa-iOS/prueba-twitter/bd"
	"github.com/GicGa-iOS/prueba-twitter/models"
)

/*UpdateProfile modify the user profile*/
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Wrong data "+err.Error(), 400)
		return
	}

	var status bool
	status, err = bd.UpdateRegister(t, IDUser)
	if err != nil {
		http.Error(w, "An error occurred while trying to modify the registry. Try again "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "The user registration has not been modified", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

package routers

import (
	"encoding/json"
	"net/http"

	"github.com/GicGa-iOS/prueba-twitter/bd"
)

/*ShowProfile allows get profile values*/
func ShowProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Must send ID parameter", http.StatusBadRequest)
		return
	}

	profile, err := bd.SearchProfile(ID)
	if err != nil {
		http.Error(w, "An error occurred while trying to search the registry", 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}

package routers

import (
	"net/http"

	"github.com/GicGa-iOS/prueba-twitter/bd"
	"github.com/GicGa-iOS/prueba-twitter/models"
)

/*RemoveFollow delete relationship between users*/
func RemoveFollow(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relation
	t.UserID = IDUser
	t.UserRelationID = ID

	status, err := bd.DeleteFollow(t)
	if err != nil {
		http.Error(w, "An error occurred while trying to delete the relationship "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "Could not delete relationship", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

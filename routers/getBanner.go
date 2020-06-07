package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/GicGa-iOS/prueba-twitter/bd"
)

/*GetBanner send banner to HTTP*/
func GetBanner(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send id parameter", http.StatusBadRequest)
		return
	}

	profile, err := bd.SearchProfile(ID)
	if err != nil {
		http.Error(w, "User not found "+err.Error(), http.StatusBadRequest)
		return
	}

	openFile, err := os.Open("uploads/banners/" + profile.Banner)
	if err != nil {
		http.Error(w, "Image not found "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, openFile)
	if err != nil {
		http.Error(w, "Error while copying image "+err.Error(), http.StatusBadRequest)
	}
}

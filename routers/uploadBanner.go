package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/GicGa-iOS/prueba-twitter/bd"
	"github.com/GicGa-iOS/prueba-twitter/models"
)

/*UploadBanner upload banner to database*/
func UploadBanner(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1]
	var filePath string = "uploads/banners/" + IDUser + "." + extension

	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error uploading image "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error copying image "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Banner = IDUser + "." + extension
	status, err = bd.UpdateRegister(user, IDUser)
	if err != nil || !status {
		http.Error(w, "Error while saving banner in database "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/GicGa-iOS/prueba-twitter/bd"
	"github.com/GicGa-iOS/prueba-twitter/models"
)

/*SaveTweet allows to save the tweet in DB*/
func SaveTweet(w http.ResponseWriter, r *http.Request) {
	var message models.SaveTweet
	err := json.NewDecoder(r.Body).Decode(&message)

	register := models.SaveTweet{
		UserID:  IDUser,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := bd.InsertTweet(register)
	if err != nil {
		http.Error(w, "An error occurred while trying to insert register, Try again "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "The tweet could not be inserted", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

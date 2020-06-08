package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GicGa-iOS/prueba-twitter/bd"
)

/*ReadFollowedTweets read followed user tweets*/
func ReadFollowedTweets(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "You must send id parameter", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "You must send the page parameter with a value greater than zero "+err.Error(), http.StatusBadRequest)
		return
	}

	response, correct := bd.GetFollowedTweets(IDUser, page)
	if !correct {
		http.Error(w, "Error reading tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

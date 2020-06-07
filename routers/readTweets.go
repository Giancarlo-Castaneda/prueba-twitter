package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GicGa-iOS/prueba-twitter/bd"
)

/*ReadTweets read the tweets*/
func ReadTweets(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 0 {
		http.Error(w, "You must send id parameter", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "You must send page parameter", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "You must send the page parameter with a value greater than zero", http.StatusBadRequest)
		return
	}

	pag := int64(page)
	response, correct := bd.ReadTweet(ID, pag)
	if !correct {
		http.Error(w, "Error reading tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
}

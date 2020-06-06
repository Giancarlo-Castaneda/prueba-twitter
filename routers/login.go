package routers

import(
	"encoding/json"
	"net/http"
	"time"
	"github.com/GicGa-iOS/prueba-twitter/bd"
	"github.com/GicGa-iOS/prueba-twitter/jwt"
	"github.com/GicGa-iOS/prueba-twitter/models"
)

/*Login execute login*/
func Login(w http.ResponseWriter, r *http.Request){
	w.Header().Add("content-type","application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil{
		http.Error(w, "User and/or password invalid "+err.Error(),400)
		return
	}

	if len(t.Email) == 0{
		http.Error(w, "User email is required",400)
		return
	}

	document, exist := bd.TryLogin(t.Email, t.Password)
	if exist == false {
		http.Error(w, "User and/or password invalid",400)
		return
	}

	jwtKey, err := jwt.CreateJWT(document)
	if err != nil{
		http.Error(w, "An error occurred while trying to generate the token "+err.Error(),400)
		return
	}

	resp := models.ResponseLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	/*-----Setting cookie from BE-----*/
	expitationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: jwtKey,
		Expires: expitationTime,
	})
}
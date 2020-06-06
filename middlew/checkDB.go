package middlew

import (
	"net/http"

	"github.com/GicGa-iOS/prueba-twitter/bd"
)

/*CheckDB is the middleware that allows to know the status of the database*/
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == false {
			http.Error(w, "Lost database connection", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}

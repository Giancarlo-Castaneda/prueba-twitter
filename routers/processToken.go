package routers

import (
	"errors"
	"strings"

	"github.com/GicGa-iOS/prueba-twitter/bd"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/GicGa-iOS/prueba-twitter/models"
)

/*Email is the email value to use on all endpoints*/
var Email string

/*IDUser is the id returned from the model, which will be used on all endpoints*/
var IDUser string

/*ProcessToken extract values from token*/
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	miKey := []byte("AprendiendoGO-paraSerMiPropioBE")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("invalid token format")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miKey, nil
	})

	if err == nil {
		_, found, _ := bd.CheckUserExists(claims.Email)
		if found == true {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, found, IDUser, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}
	return claims, false, string(""), err
}

package jwt

import (
	"time"

	"github.com/GicGa-iOS/prueba-twitter/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*CreateJWT creates the encrypted JWT*/
func CreateJWT(t models.User) (string, error) {
	myKey := []byte("AprendiendoGO-paraSerMiPropioBE")

	payload := jwt.MapClaims{
		"email":     t.Email,
		"name":      t.Name,
		"lastName":  t.LastName,
		"birthDate": t.BirthDate,
		"biography": t.Biography,
		"location":  t.Location,
		"webSite":   t.WebSite,
		"_id":       t.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}

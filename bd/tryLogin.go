package bd

import (
	"github.com/GicGa-iOS/prueba-twitter/models"
	"golang.org/x/crypto/bcrypt"
)

/*TryLogin check login to DB*/
func TryLogin(email string, password string) (models.User, bool) {
	usu, found, _ := CheckUserExists(email)
	if found == false {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}

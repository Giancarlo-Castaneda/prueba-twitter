package models

/*ResponseLogin contains the token that is returned with the login*/
type ResponseLogin struct {
	Token string `json:"token,omitempty"`
}

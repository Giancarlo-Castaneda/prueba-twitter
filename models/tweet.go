package models

/*Tweet capture of the body, the message we received*/
type Tweet struct {
	Messasge string `bson:"message" json:"message"`
}

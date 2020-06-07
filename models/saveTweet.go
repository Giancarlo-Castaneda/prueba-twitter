package models

import "time"

/*SaveTweet is the model of the tweet in the db*/
type SaveTweet struct {
	UserID  string    `bson:"userID" json:"userID,omitempty"`
	Message string    `bson:"message" json:"message,omitempty"`
	Date    time.Time `bson:"date" json:"date,omitempty"`
}

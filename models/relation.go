package models

/*Relation model used to save the relationship between users*/
type Relation struct {
	UserID         string `bson:"userid" json:"userId"`
	UserRelationID string `bson:"userrelationid" json:"userRelationId"`
}

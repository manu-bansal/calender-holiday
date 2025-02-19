package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Holiday struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Date        string             `bson:"date" json:"date"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
}

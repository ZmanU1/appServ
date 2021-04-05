package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName string             `bson:"firstname,omitempty" json:"firstname,omitempty"`
	LastName  string             `bson:"lastname,omitempty" json:"lastname,omitempty"`
	Age       int                `bson:"age,omitempty" json:"age,omitempty"`
}

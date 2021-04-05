package model

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// You will be using this Trainer type later in the program
type Trainer struct {
	Id   primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name string             `bson:"name"`
	Age  int                `bson:"age"`
	City string             `bson:"city"`
}

func DescribeCollection(collection []*Trainer) {
	for i := range collection {
		fmt.Printf("%s is from %s and is %d years old\n", collection[i].Name, collection[i].City, collection[i].Age)
	}
}
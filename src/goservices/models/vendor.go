package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Vendor struct {
	Id       primitive.ObjectID `json:"id,omitempty"  bson:"_id,omitempty`
	Name     string             `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	Location string             `json:"location,omitempty" bson:"location,omitempty" validate:"required"`
	Title    string             `json:"title,omitempty" bson:"title,omitempty" validate:"required"`
}

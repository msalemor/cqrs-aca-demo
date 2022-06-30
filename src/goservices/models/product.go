package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	Price       float32            `json:"price,omitempty" bson:"price,omitempty" validate:"required"`
	Weight      float32            `json:"weight,omitempty" bson:"weight,omitempty" validate:"required"`
	Size        string
	CreatedDate time.Time `json:"createdDate" bson:"createdDate"`
}

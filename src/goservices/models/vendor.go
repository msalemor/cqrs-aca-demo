package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Vendor struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	VendorCode  string             `json:"vendorCode,omitempty" bson:"vendorCode,omitempty" validate:"required"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	Contact     string             `json:"contact,omitempty" bson:"contact,omitempty" validate:"required"`
	Phone       string             `json:"phone,omitempty" bson:"phone,omitempty" validate:"required"`
	Email       string             `json:"email,omitempty" bson:"email,omitempty" validate:"required"`
	CreatedDate time.Time          `json:"createdDate" bson:"createdDate"`
}

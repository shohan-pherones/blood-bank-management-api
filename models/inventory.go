package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Inventory struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	BloodType  BloodType          `bson:"blood_type" json:"blood_type"`
	Quantity   int                `bson:"quantity" json:"quantity"`
	ExpiryDate string             `bson:"expiry_date" json:"expiry_date"`
	ReceivedAt primitive.DateTime `bson:"received_at" json:"received_at"`
	UpdatedAt  primitive.DateTime `bson:"updated_at" json:"updated_at"`
}

package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRole string

const (
	Admin     UserRole = "admin"
	Donor     UserRole = "donor"
	Recipient UserRole = "recipient"
)

type UserModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName string             `bson:"first_name" json:"first_name"`
	LastName  string             `bson:"last_name" json:"last_name"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	Phone     string             `bson:"phone" json:"phone"`
	Address   string             `bson:"address" json:"address"`
	Role      UserRole           `bson:"role" json:"role"`
	CreatedAt primitive.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at" json:"updated_at"`
}

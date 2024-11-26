package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BloodType string

const (
	APositive  BloodType = "A+"
	ANegative  BloodType = "A-"
	BPositive  BloodType = "B+"
	BNegative  BloodType = "B-"
	ABPositive BloodType = "AB+"
	ABNegative BloodType = "AB-"
	OPositive  BloodType = "O+"
	ONegative  BloodType = "O-"
)

type Sex string

const (
	Male   Sex = "male"
	Female Sex = "female"
	Other  Sex = "other"
)

type DonorModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	FirstName string             `bson:"first_name" json:"first_name"`
	LastName  string             `bson:"last_name" json:"last_name"`
	Email     string             `bson:"email" json:"email"`
	Phone     string             `bson:"phone" json:"phone"`
	Address   string             `bson:"address" json:"address"`
	BloodType BloodType          `bson:"blood_type" json:"blood_type"`
	Age       int                `bson:"age" json:"age"`
	Sex       Sex                `bson:"sex" json:"sex"`
	Donations int                `bson:"donations" json:"donations"`
	CreatedAt primitive.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at" json:"updated_at"`
}

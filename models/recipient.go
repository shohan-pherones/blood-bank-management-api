package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransfusionHistory struct {
	DonorID   primitive.ObjectID `bson:"donor_id" json:"donor_id"`
	Date      primitive.DateTime `bson:"date" json:"date"`
	BloodType BloodType          `bson:"blood_type" json:"blood_type"`
}

type RecipientModel struct {
	ID                 primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	UserID             primitive.ObjectID   `bson:"user_id" json:"user_id"`
	FirstName          string               `bson:"first_name" json:"first_name"`
	LastName           string               `bson:"last_name" json:"last_name"`
	Email              string               `bson:"email" json:"email"`
	Phone              string               `bson:"phone" json:"phone"`
	Address            string               `bson:"address" json:"address"`
	BloodType          BloodType            `bson:"blood_type" json:"blood_type"`
	Age                int                  `bson:"age" json:"age"`
	Sex                Sex                  `bson:"sex" json:"sex"`
	TransfusionHistory []TransfusionHistory `bson:"transfusion_history" json:"transfusion_history"`
	CreatedAt          primitive.DateTime   `bson:"created_at" json:"created_at"`
	UpdatedAt          primitive.DateTime   `bson:"updated_at" json:"updated_at"`
}

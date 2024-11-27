package services

import (
	"context"
	"errors"
	"time"

	"github.com/shohan-pherones/blood-bank-management.git/database"
	"github.com/shohan-pherones/blood-bank-management.git/models"
	"github.com/shohan-pherones/blood-bank-management.git/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct{}

func (us *UserService) RegisterUser(user *models.UserModel) error {
	if !utils.ValidateEmail(user.Email) {
		return errors.New("invalid email format")
	}

	existingUser := database.UserColl.FindOne(context.TODO(), bson.M{"email": user.Email})
	if existingUser.Err() == nil {
		return errors.New("email already in use")
	}

	if !utils.ValidatePassword(user.Password) {
		return errors.New("password must be at least 8 characters long, contain an uppercase letter, a number, and a special character")
	}

	if !utils.ValidateName(user.FirstName) || !utils.ValidateName(user.LastName) {
		return errors.New("names can only contain letters and spaces")
	}

	if !utils.ValidatePhone(user.Phone) {
		return errors.New("invalid phone number format")
	}

	if err := utils.ValidateRole(string(user.Role)); err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	user.ID = primitive.NewObjectID()
	user.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	user.UpdatedAt = user.CreatedAt

	_, err = database.UserColl.InsertOne(context.TODO(), user)
	if err != nil {
		return errors.New("failed to save user to the database")
	}

	return nil
}

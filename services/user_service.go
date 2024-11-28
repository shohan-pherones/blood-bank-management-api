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
	"go.mongodb.org/mongo-driver/mongo"
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

func (us *UserService) LoginUser(email, password string) (*models.UserModel, error) {
	if !utils.ValidateEmail(email) {
		return nil, errors.New("invalid email format")
	}

	var user models.UserModel
	err := database.UserColl.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("failed to query the database")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid password")
	}

	return &user, nil
}

func (us *UserService) GetUsers() ([]models.UserModel, error) {
	var users []models.UserModel
	cursor, err := database.UserColl.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, errors.New("failed to query the database")
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var user models.UserModel
		if err := cursor.Decode(&user); err != nil {
			return nil, errors.New("failed to decode user")
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, errors.New("error while iterating over database cursor")
	}

	return users, nil
}

func (us *UserService) GetUser(userID string) (*models.UserModel, error) {
	var user models.UserModel

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	err = database.UserColl.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("failed to query the database")
	}

	return &user, nil
}

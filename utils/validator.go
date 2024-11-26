package utils

import (
	"errors"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/go-playground/validator/v10"
	"github.com/shohan-pherones/blood-bank-management.git/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ValidateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func ValidatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasUpper := false
	hasNumber := false
	hasSpecial := false

	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUpper = true
		}
		if unicode.IsDigit(char) {
			hasNumber = true
		}
		if strings.ContainsAny(string(char), "!@#$%^&*()-_=+[]{}|;:,.<>?") {
			hasSpecial = true
		}
	}

	return hasUpper && hasNumber && hasSpecial
}

func ValidateName(name string) bool {
	re := regexp.MustCompile(`^[a-zA-Z\s]+$`)
	return re.MatchString(name)
}

func ValidatePhone(phone string) bool {
	re := regexp.MustCompile(`^\d{10}$`)
	return re.MatchString(phone)
}

func ValidateAddress(address string) bool {
	return len(address) > 0
}

func ValidateBloodType(bloodType string) bool {
	validBloodTypes := []string{"A+", "A-", "B+", "B-", "AB+", "AB-", "O+", "O-", "O", "AB"}

	bloodType = strings.ToUpper(strings.TrimSpace(bloodType))

	for _, validType := range validBloodTypes {
		if bloodType == validType {
			return true
		}
	}
	return false
}

func ValidateAge(fl validator.FieldLevel) bool {
	birthdate, ok := fl.Field().Interface().(time.Time)
	if !ok {
		return false
	}

	currentDate := time.Now()
	age := currentDate.Year() - birthdate.Year()
	if currentDate.Before(birthdate.AddDate(age, 0, 0)) {
		age--
	}

	return age >= 18 && age <= 100
}

func ValidateSex(fl validator.FieldLevel) bool {
	sex := fl.Field().String()
	allowedValues := []string{"male", "female", "other"}

	sex = strings.ToLower(sex)
	for _, allowed := range allowedValues {
		if sex == allowed {
			return true
		}
	}
	return false
}

func ValidateDonations(donations int) error {
	if donations < 0 {
		return errors.New("donations must be a non-negative integer")
	}
	return nil
}

func ValidateQuantity(quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity must be a positive integer")
	}
	return nil
}

func ValidateExpiryDate(expiryDate time.Time) error {
	if expiryDate.Before(time.Now()) {
		return errors.New("expiry date must be a future date")
	}
	return nil
}

var validRoles = []string{"admin", "user"}

func ValidateRole(role string) error {
	role = strings.ToLower(role)
	for _, validRole := range validRoles {
		if role == validRole {
			return nil
		}
	}
	return errors.New("invalid role: must be 'admin' or 'user'")
}

func ValidateTransfusionHistory(transfusionHistory models.TransfusionHistory) error {
	if transfusionHistory.DonorID == primitive.NilObjectID {
		return errors.New("invalid donor ID")
	}

	if transfusionHistory.Date.Time().After(time.Now()) {
		return errors.New("transfusion date cannot be in the future")
	}

	validBloodTypes := map[models.BloodType]bool{
		models.APositive:  true,
		models.ANegative:  true,
		models.BPositive:  true,
		models.BNegative:  true,
		models.ABPositive: true,
		models.ABNegative: true,
		models.OPositive:  true,
		models.ONegative:  true,
	}

	if _, valid := validBloodTypes[transfusionHistory.BloodType]; !valid {
		return errors.New("invalid blood type")
	}

	return nil
}

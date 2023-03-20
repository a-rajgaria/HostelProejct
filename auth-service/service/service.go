package service

import (
	"errors"
	"os"
	"time"

	"github.com/a-rajgaria/HostelProject/models"
	"github.com/a-rajgaria/HostelProject/repository"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func RegisterCustomer(data map[string]string) models.Customer {

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	customer := models.Customer{
		Name: data["name"],
		Email: data["email"],
		Mobile: data["mobile"],
		Password: string(password),
	}

	return repository.DBCreate(customer)

}

func LoginCustomer(data map[string]string) (string, error) {
	customer := repository.DBGetByEmail(data["email"])
	
	if customer.Id == ""{
		return "", errors.New("email not found")
	}

	if err:= bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(data["password"])); err != nil{
		return "", errors.New("incorrect password")
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"customerId": customer.Id,
		"expiresAt": time.Now().Add(time.Minute*1).Unix(),
	})

	token, err := claims.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil{
		return "", errors.New("could not login")
	}
	return token, nil
}
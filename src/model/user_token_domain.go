package model

import (
	"fmt"
	"os"
	"time"

	"github.com/LucasKeley/CRUD_Go/src/configuration/rest_err"
	"github.com/golang-jwt/jwt"
)

var (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

func (ud *userDomain) GenerateToken() (string, *rest_err.RestErr) {
	secret := os.Getenv("JWT_SECRET_KEY")

	claims := jwt.MapClaims{
		"id":    ud.id,
		"email": ud.email,
		"name":  ud.name,
		"age":   ud.age,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenSring, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", rest_err.NewInternalServerError(
			fmt.Sprintf("error trying to generate jwt token, err=%s", err.Error()))
	}

	return tokenSring, nil
}

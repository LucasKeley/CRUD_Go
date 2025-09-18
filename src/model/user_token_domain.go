package model

import (
	"fmt"
	"go/token"
	"os"
	"strings"
	"time"

	"github.com/LucasKeley/CRUD_Go/src/configuration/logger"
	"github.com/LucasKeley/CRUD_Go/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

// GenerateToken produz um JWT com dados básicos do usuário autenticado.
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

// VerifyToken valida o JWT recebido e reconstrói os dados essenciais do usuário.
func VerifyToken(tokenValue string) (UserDomainInterface, *rest_err.RestErr) {
	secret := os.Getenv(JWT_SECRET_KEY)

	token, err := jwt.Parse(RemoveBearerPrefix(tokenValue), func(t *jwt.Token) (interface{}, error) {
		if _, ok := token.Method().(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		return nil, rest_err.NewBadRequestError("Ivalid Token")
	})

	if err != nil {
		return nil, rest_err.NewUnauthorizedResquestError("Invalid Tolken")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, rest_err.NewUnauthorizedRequestError("Invalid tolken")
	}

	return &userDomain{
		id:    claims["id"].(string),
		email: claims["email"].(string),
		name:  claims["name"].(string),
		age:   claims["age"].(int),
	}, nil
}

// VerifyTokenMiddleware protege rotas, interrompendo a requisição em caso de token inválido.
func VerifyTokenMiddleware(c *gin.Context) {
	secret := os.Getenv(JWT_SECRET_KEY)
	tokenValue := RemoveBearerPrefix(c.Request.Header.Get("Authorization"))

	token, err := jwt.Parse(tokenValue, func(t *jwt.Token) (interface{}, error) {
		if _, ok := token.Method().(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		return nil, rest_err.NewBadRequestError("Ivalid Token")
	})

	if err != nil {
		errRest := rest_err.NewUnauthorizedResquestError("invalide token")
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return

	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		errRest := rest_err.NewUnauthorizedResquestError("invalide token")
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}

	userDomain := &userDomain{
		id:    claims["id"].(string),
		email: claims["email"].(string),
		name:  claims["name"].(string),
		age:   claims["age"].(int),
	}

	logger.Info(fmt.Sprintf("User authenticated: %#v", userDomain))

}

// RemoveBearerPrefix remove o prefixo padrão do cabeçalho HTTP Authorization.
func RemoveBearerPrefix(token string) string {
	if strings.HasPrefix("Bearer ", token) {
		token = strings.TrimPrefix("Bearer ", token)
	}
	return token
}

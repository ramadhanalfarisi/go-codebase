package helpers

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/ramadhanalfarisi/go-codebase-kocak/models"
)

var APPLICATION_NAME = "belajar-golang"
var LOGIN_EXPIRATION_DURATION = time.Duration(168) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("belajargolangkocak")

type UserClaims struct {
	jwt.StandardClaims
	UserId    uuid.UUID `json:"userId"`
	UserEmail string    `json:"userEmail"`
	UserRole  string    `json:"userRole"`
}

func GenerateJWT(userData models.User) string{
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer: APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
		UserId: userData.Id,
		UserEmail: userData.Email,
		UserRole: userData.Role,
	}
	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claims)
	signedToken, _ := token.SignedString(JWT_SIGNATURE_KEY)
	return signedToken
}

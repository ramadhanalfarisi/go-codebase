package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserDetail struct {
	Id    int `json:"id"`
	Email string `json:"email"`
	Roles string `json:"roles"`
}

func GenerateJWT(data UserDetail, expiryTime *time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"iat":    jwt.NewNumericDate(time.Now()),
		"userId": data.Id,
		"email":  data.Email,
		"roles":  data.Roles,
	}
	if expiryTime != nil {
		claims["exp"] = jwt.NewNumericDate(time.Now().Add(*expiryTime))
	} else {
		claims["exp"] = jwt.NewNumericDate(time.Now().Add(30 * time.Minute))
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := []byte(os.Getenv("JWT_SECRET"))

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseUserJWT(tokenString string) (jwt.MapClaims, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenMalformed
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrTokenInvalidClaims
}

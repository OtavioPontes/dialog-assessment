package authentication

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/otaviopontes/api-go/src/config"
)

func CreateToken(userId uuid.UUID) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).UnixMilli()
	permissions["userId"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.SecretKey))
}

func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)

	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("invalid token")
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}

func ExtractUserId(r *http.Request) (uuid.UUID, error) {
	tokenString := extractToken(r)

	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return uuid.Nil, err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, err := uuid.Parse(fmt.Sprintf("%s", permissions["userId"]))
		if err != nil {
			return uuid.Nil, err
		}
		return userId, nil
	}
	return uuid.Nil, errors.New("invalid token")
}

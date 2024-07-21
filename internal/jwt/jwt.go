package jwt

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
)

type contextKey string

const UserIDKey contextKey = "user_id"

type Claims struct {
	UserID uint32 `json:"user_id"`
	jwt.StandardClaims
}

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func init() {
	if len(jwtKey) == 0 {
		log.Fatal("JWT_SECRET_KEY environment variable is not set")
	}
}

func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}

func NewContextWithUserID(ctx context.Context, userID uint32) context.Context {
	return context.WithValue(ctx, UserIDKey, userID)
}

func GetUserID(ctx context.Context) (uint32, bool) {
	id, ok := ctx.Value(UserIDKey).(uint32)
	return id, ok
}

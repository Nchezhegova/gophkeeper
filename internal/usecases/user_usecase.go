package usecases

import (
	"errors"
	"github.com/Nchezhegova/gophkeeper/internal/entities"
	"github.com/Nchezhegova/gophkeeper/internal/interfaces/repository"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var jwtKey = []byte("your_secret_key")

type UserUseCase struct {
	UserRepository repository.UserRepository
}

type Claims struct {
	UserID uint32 `json:"user_id"`
	jwt.StandardClaims
}

func (uc *UserUseCase) Register(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := entities.User{Username: username, Password: string(hashedPassword)}
	return uc.UserRepository.Create(user)
}

func (uc *UserUseCase) Login(username, password string) (string, error) {
	user, err := uc.UserRepository.GetByUsername(username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: uint32(user.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (uc *UserUseCase) GetUserByUsername(username string) (entities.User, error) {
	return uc.UserRepository.GetByUsername(username)
}

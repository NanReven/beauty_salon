package service

import (
	"beauty_salon/internal/adapter/repository"
	"beauty_salon/internal/domain/entity"
	"crypto/sha1"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	UserId int    `json:"id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (serv *UserService) Register(input *entity.User) (int, error) {
	input.Password = generatePasswordHash(input.Password)
	return serv.repo.CreateUser(input)
}

func (serv *UserService) GenerateToken(email, password string) (string, error) {
	user, err := serv.repo.GetUser(email, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserClaims{
		UserId: user.Id,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
		},
	})
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (serv *UserService) ParseToken(token string) (int, string, error) {
	authToken, err := jwt.ParseWithClaims(token, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return 0, "", err
	}

	claims, ok := authToken.Claims.(*UserClaims)
	if !ok {
		return 0, "", errors.New("invalid token claims")
	}
	return claims.UserId, claims.Role, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("PASSWORD_SALT"))))
}

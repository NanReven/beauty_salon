package usecase

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
	UserId   int  `json:"id"`
	IsMaster bool `json:"is_master"`
	jwt.RegisteredClaims
}

type UserUsecase struct {
	repo repository.User
}

func NewUserUsecase(repo repository.User) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (uc *UserUsecase) Register(input *entity.User) (int, error) {
	input.Password = generatePasswordHash(input.Password)
	return uc.repo.CreateUser(input)
}

func (uc *UserUsecase) GenerateToken(email, password string) (string, error) {
	user, err := uc.repo.GetUser(email, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserClaims{
		UserId:   user.Id,
		IsMaster: user.IsMaster,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
		},
	})
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (uc *Usecase) ParseToken(token string) (int, bool, error) {
	authToken, err := jwt.ParseWithClaims(token, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return 0, false, err
	}

	claims, ok := authToken.Claims.(*UserClaims)
	if !ok {
		return 0, false, errors.New("invalid token claims")
	}
	return claims.UserId, claims.IsMaster, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("PASSWORD_SALT"))))
}

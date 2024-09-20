package usecase

import (
	"beauty_salon/internal/adapter/repository"
	"beauty_salon/internal/domain/entity"
	"crypto/sha1"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

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

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("PASSWORD_SALT"))))
}

type UserClaims struct {
	UserId   int  `json:"id"`
	IsMaster bool `json:"is_master"`
	jwt.RegisteredClaims
}

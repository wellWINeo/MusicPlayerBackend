package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	_ "github.com/sirupsen/logrus"
	"github.com/wellWINeo/MusicPlayerBackend"
	"github.com/wellWINeo/MusicPlayerBackend/pkg/repository"
)


const (
	salt = "asdfj324eo5kj435kj321aj"
	tokenTTL = 12 * time.Hour
	tokenSecret = "sdf734bjhrb34l673hj32"
)


type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) CreateUser(user MusicPlayerBackend.User) (int, error) {
	user.Password = s.GenerateHashPassword(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateHashPassword(passwd string) string {
	hash := sha1.New()
	hash.Write([]byte(passwd))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, s.GenerateHashPassword(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt: time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(tokenSecret))
}

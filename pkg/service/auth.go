package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"math/rand"
	"net/smtp"
	"regexp"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	_ "github.com/sirupsen/logrus"
	"github.com/wellWINeo/MusicPlayerBackend"
	"github.com/wellWINeo/MusicPlayerBackend/pkg/repository"
)

const (
	salt        = "asdfj324eo5kj435kj321aj"
	tokenTTL    = 12 * time.Hour
	tokenSecret = "sdf734bjhrb34l673hj32"
	regexMail = `^([a-z0-9_-]+\.)*[a-z0-9_-]+@[a-z0-9_-]+(\.[a-z0-9_-]+)*\.[a-z]{2,6}$`
	regexOnlyASCII = `^[\x00-\x7F]*$`
)

var (
	EmptyPassword = errors.New("Empty password field")
	InvalidCharsInPassword = errors.New("Invalid chars in password")
	EmptyUsername = errors.New("Empty username field")
	InvalidCharsInUsername = errors.New("Invalid chars in username")
	NotValidMail = errors.New("Invalid email field")
)

type AuthService struct {
	verificationCodes map[int]MusicPlayerBackend.User
	mailConfig        MailConfig
	repo              repository.Authorization
}
type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type MailConfig struct {
	Port int
	Host,
	MailBox,
	Password string
}

func NewAuthService(repo repository.Authorization, mailConfig MailConfig) *AuthService {

	return &AuthService{
		verificationCodes: make(map[int]MusicPlayerBackend.User),
		repo:              repo,
		mailConfig:        mailConfig,
	}
}

func (s *AuthService) ValidateUser(user MusicPlayerBackend.User) error {
	if user.Username == "" {
		return EmptyUsername
	}

	matched, err := regexp.MatchString(regexMail, user.Email)
	if !matched || err != nil {
		return NotValidMail
	}

	matched, err = regexp.MatchString(regexOnlyASCII, user.Username)
	if !matched || err != nil {
		return InvalidCharsInUsername
	}

	if user.Password == "" {
		return EmptyPassword
	}

	matched, err = regexp.MatchString(regexOnlyASCII, user.Password)
	if !matched || err != nil {
		return InvalidCharsInPassword
	}

	return nil
}

func (s *AuthService) CreateUser(user MusicPlayerBackend.User) (int, error) {
	user.Password = s.GenerateHashPassword(user.Password)
	if err := s.ValidateUser(user); err != nil {
		return 0, err
	}

	id, err := s.repo.CreateUser(user)
	if err != nil {
		return id, err
	}

	if user.Referal != 0 {
		if err := s.repo.CreateReferal(user.Referal, id); err != nil {
			return id, err
		}
	}

	return id, s.SendCode(user)
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
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(tokenSecret))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return 0, nil
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("can't parse token claims")
	}
	return claims.UserId, nil
}

// email verification
func (s *AuthService) SendCode(user MusicPlayerBackend.User) error {
	var code int
	rand.Seed(time.Now().Unix())
	ok := true
	for ok {
		code = int(rand.Int31n(10000))
		_, ok = s.verificationCodes[int(code)]
	}
	s.verificationCodes[int(code)] = user

	to := []string{user.Email}
	addr := fmt.Sprintf("%s:%d", s.mailConfig.Host, s.mailConfig.Port)

	msg := []byte(fmt.Sprintf("To: %s\r\n"+
		"Subject: Verification code\r\n"+
		"Code - %d\r\n",
		user.Email, code))
	auth := smtp.PlainAuth("", s.mailConfig.MailBox, s.mailConfig.Password,
		s.mailConfig.Host)
	err := smtp.SendMail(addr, auth, s.mailConfig.MailBox, to, msg)

	return err
}

func (s *AuthService) Verify(code int) (MusicPlayerBackend.User, bool) {
	value, ok := s.verificationCodes[code]
	return value, ok
}

func (s *AuthService) UpdateUser(user MusicPlayerBackend.User) error {
	// user validation
	err := s.ValidateUser(user)

	if err != nil && err != EmptyPassword {
		return err
	}

	if err == EmptyPassword {
		user.Password = s.GenerateHashPassword(user.Password)
	}
	return s.repo.UpdateUser(user)
}

func (s *AuthService) GetUser(id int) (MusicPlayerBackend.User, error) {
	return s.repo.GetUserById(id)
}

func (s *AuthService) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}

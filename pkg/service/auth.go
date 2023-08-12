package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/ShatAlex/chat"
	"github.com/ShatAlex/chat/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "dfW52sz091dfAPLmGgZ7"
	signingKey = "asd%awASBd#as#LtN#sad124"
	tokeTTL    = 8 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

type AuthService struct {
	rep repository.Authorization
}

func NewAuthService(rep repository.Authorization) *AuthService {
	return &AuthService{rep: rep}
}

func (s *AuthService) CreateUser(user chat.User) (int, error) {

	user.Password = generatePasswordHash(user.Password)
	return s.rep.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user_id, err := s.rep.GetUserId(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokeTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserID: user_id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims aren't of type *tokenCLaims")
	}
	return claims.UserID, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

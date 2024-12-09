package services

import (
	"crypto/sha1"
	"errors"
	"fmt"
	minichatgo "mini_chat_go"
	"mini_chat_go/pkg/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "dsfgrskdgfkdsfslsdf"
	tokenTTL   = 12 * time.Hour
	signingKey = "grkjk#4#%35FSFJlja4353KSFjH"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}
type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (as *AuthService) CreateUser(user minichatgo.User) (int, error) {
	user.Password = generatePassHash(user.Password)
	return as.repo.CreateUser(user)
}

func (as *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := as.repo.GetUser(username, generatePassHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		}, user.ID,
	})
	return token.SignedString([]byte(signingKey))
}

func (as *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, fmt.Errorf("invalid token parsing method %s", err.Error())
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not type of *tokenClaims")
	}
	return claims.UserID, nil
}

func generatePassHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}



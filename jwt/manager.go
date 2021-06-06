package jwt

import (
	"fmt"
	"github.com/PulseyTeam/game-server/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	secretKey     = "0pPeA3tn0BUesH9dtjptZsZpuYHZtCDFqUFH7EdiVw4U8APE6uTNg53LXqq1EAa"
	tokenDuration = 4 * time.Hour
)

type Manager struct {
	secretKey     string
	tokenDuration time.Duration
}

type UserClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func NewManager() *Manager {
	return &Manager{secretKey, tokenDuration}
}

func (manager *Manager) Generate(user *model.User) (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.tokenDuration).Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   user.ID.Hex(),
		},
		Username: user.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.secretKey))
}

func (manager *Manager) Verify(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return []byte(manager.secretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

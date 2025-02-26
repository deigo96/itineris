package service

import (
	"time"

	"github.com/deigo96/itineris/app/config"
	"github.com/deigo96/itineris/app/internal/model"
	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	GenerateToken(customClaim model.CustomClaims) (string, error)
	// ValidateToken(token string) (string, error)
}

type jwtService struct {
	config *config.Config
}

func NewJWTService(config *config.Config) JWTService {
	return &jwtService{
		config: config,
	}
}

func (s *jwtService) GenerateToken(customClaim model.CustomClaims) (string, error) {
	jwtClaims := jwt.MapClaims{
		"id":   customClaim.ID,
		"nip":  customClaim.Nip,
		"role": customClaim.Role,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	tokenString, err := token.SignedString([]byte(s.config.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

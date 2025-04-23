package service

import (
    "go-api/config"
    "time"
    "github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
    GenerateToken(userId int, email string, role string) (string, error)
    ValidateToken(tokenString string) (*jwt.Token, error)
}

type jwtService struct {
    secretKey string
    issuer    string
}

func NewJWTService() JWTService {
    return &jwtService{
        secretKey: config.SecretKey,
        issuer:    "go-api",
    }
}

func (j *jwtService) GenerateToken(userId int, email string, role string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userId,
        "email":   email,
        "role":    role,
        "exp":     time.Now().Add(time.Hour * time.Duration(config.TokenExpiration)).Unix(),
        "iat":     time.Now().Unix(),
        "iss":     j.issuer,
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(j.secretKey))
}

func (j *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, jwt.ErrSignatureInvalid
        }
        return []byte(j.secretKey), nil
    })
}
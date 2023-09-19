package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Config struct {
	Secret string        `yaml:"secret"`
	Expire time.Duration `yaml:"expire"`
}

type CustomClaims struct {
	// 额外的信息
	uid int64
	jwt.RegisteredClaims
}

func SignJwtToken(uid int64, config Config) (string, error) {
	claims := &CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Hissssss",
			Subject:   "json web token",
			Audience:  nil,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.Expire)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.NewV4().String(),
		},
		uid: uid,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return "", err
	}
	return signedString, nil
}

func ParseJwtToken(sign string, config Config) (*CustomClaims, error) {
	parseToken, err := jwt.ParseWithClaims(sign, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := parseToken.Claims.(*CustomClaims); ok && parseToken.Valid {
		return claims, nil
	}
	return nil, jwt.ErrSignatureInvalid
}

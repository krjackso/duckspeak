package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	jwtAccessExpire   = 60 * time.Minute
	jwtAccessAudience = "self"
)

type IAuthenticator interface {
	NewAccessToken(string) (string, int64)
	VerifyAccessToken(string) (string, bool)
}

type Authenticator struct {
	issuer string
	secret string
}

func NewAuthenticator(issuer, secret string) *Authenticator {
	return &Authenticator{
		issuer: issuer,
		secret: secret,
	}
}

func (self *Authenticator) NewAccessToken(deviceId string) (token string, expiresIn int64) {
	expiresAt := time.Now().Add(jwtAccessExpire).Unix()

	claims := &jwt.StandardClaims{
		Issuer:    self.issuer,
		IssuedAt:  time.Now().Unix(),
		Audience:  jwtAccessAudience,
		Subject:   deviceId,
		ExpiresAt: expiresAt,
	}

	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwt.SignedString([]byte(self.secret))
	if err != nil {
		panic(err)
	}

	return token, int64(jwtAccessExpire.Seconds())
}

func (self *Authenticator) VerifyAccessToken(tokenString string) (deviceId string, ok bool) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(self.secret), nil
	})

	if err != nil {
		return "", false
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok || !token.Valid {
		return "", false
	}

	if !claims.VerifyIssuer(self.issuer, true) {
		return "", false
	}

	if !claims.VerifyAudience(jwtAccessAudience, true) {
		return "", false
	}

	deviceId = claims.Subject

	return deviceId, true
}

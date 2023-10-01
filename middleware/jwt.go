package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	key        []byte
	signMethod jwt.SigningMethod
}

func NewJWT(key []byte, sign jwt.SigningMethod) *JWT {
	return &JWT{key, sign}
}

func (j *JWT) GenerateToken(user string) (string, error) {
	claims := jwt.MapClaims{
		"iss": "my-auth-server",
		"sub": user,
	}

	token := jwt.NewWithClaims(j.signMethod, claims)

	return token.SignedString(j.key)
}

func (j *JWT) ValidateToken(token string) (err error) {
	tokenJwt, err := j.Parse(token)
	if err != nil {
		return fmt.Errorf("error parsing token: %v", err)
	}
	if _, ok := tokenJwt.Claims.(jwt.MapClaims); !ok && !tokenJwt.Valid {
		return fmt.Errorf("token not valid")
	}
	return
}

func (j *JWT) Parse(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return j.key, nil
	})
}

func (j *JWT) GetUsername(token string) (data string, err error) {
	tokenJwt, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return j.key, nil
	})
	if err != nil {
		return "", fmt.Errorf("error parsing token: %v", err)
	}

	claims, ok := tokenJwt.Claims.(jwt.MapClaims)
	if !ok && !tokenJwt.Valid {
		return "", fmt.Errorf("token not valid")
	}

	data, err = claims.GetSubject()
	return

}

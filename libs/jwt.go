package libs

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	Key string
}

var (
	token *jwt.Token
)

func NewJWT(key string) *JWT {
	token = jwt.New(jwt.SigningMethodHS512)
	return &JWT{
		Key: key,
	}
}

func (j JWT) Encode() (string, error) {
	tokenString, err := token.SignedString([]byte(j.Key))
	return tokenString, err
}

func (j JWT) AddDataString(key, value string) {
	token.Claims.(jwt.MapClaims)[key] = value
}

func (j JWT) AddDataInt(key string, value int) {
	token.Claims.(jwt.MapClaims)[key] = value
}

func (j JWT) AddExpiredDate(value int) {
	token.Claims.(jwt.MapClaims)["exp"] = time.Now().Add(time.Duration(value) * time.Hour).Unix()
	token.Claims.(jwt.MapClaims)["iat"] = time.Now().Unix()
}

func (j JWT) IsExpired(exp float64) bool {
	if time.Now().Unix() > int64(exp) {
		return false
	}
	return true
}

func (j JWT) isValid(token string) bool {
	if tkn, claims, err := j.Decode(token); err != nil && tkn.Valid {
		if !j.IsExpired(claims["exp"].(float64)) {
			return true
		}
		return false
	}
	return false
}

func (j JWT) Decode(t string) (*jwt.Token, jwt.MapClaims, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Key), nil
	})
	return token, token.Claims.(jwt.MapClaims), err
}

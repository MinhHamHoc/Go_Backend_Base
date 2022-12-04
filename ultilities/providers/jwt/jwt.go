package jwt

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService struct {
	JwtKey string
}

var expirationTime = time.Now().Add(time.Minute * 10)

type Claims struct {
	Data interface{} `json:"data"`
	jwt.StandardClaims
}

func (JS JWTService) IssueJWT(data interface{}) (string, error) {
	claims := Claims{
		Data: data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(JS.JwtKey))
}
func (JS JWTService) ValidateJWT(tokenString string, data interface{}) error {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(JS.JwtKey), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return err
		}
	}
	if !token.Valid {
		return errors.New("token not valid")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return JS.parseDataFromMap(claims, data)
	}

	return fmt.Errorf("invalid JWTService")
}
func (JS JWTService) parseDataFromMap(m jwt.MapClaims, out interface{}) error {
	data, ok := m["data"]
	if !ok || data == nil {
		return fmt.Errorf("invalid JWTService Claims: Data")
	}

	buffers, e := json.Marshal(data)
	if e != nil {
		return e
	}

	return json.Unmarshal(buffers, out)
}
func NewJWT(Key string) *JWTService {
	service := new(JWTService)
	service.JwtKey = Key
	return service
}

package api

import (
	"backendbase/models"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type AuthHandler struct {
	JWTSecret string
}

type AuthLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	jwt.StandardClaims
}

func (h *AuthHandler) makeClaims(account models.Account) Claims {
	return Claims{
		StandardClaims: jwt.StandardClaims{
			Id:        bson.NewObjectId().Hex(),
			Audience:  string(account.ID),
			Subject:   account.Email,
			ExpiresAt: time.Now().AddDate(10, 0, 0).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
}

func (h *AuthHandler) CreateToken(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(h.JWTSecret))
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var data AuthLogin
	if err := BindJSON(r, data); err != nil {
		WriteJSON(w, http.StatusRequestTimeout, ResponseBody{
			Message: err.Error(),
		})
		return
	}
	if data.Email == "" || data.Password == "" {
		WriteJSON(w, http.StatusBadRequest, ResponseBody{
			Message: "email or password is empty",
			Code:    http.StatusBadRequest,
		})
		return
	}
	account, err := h.AccountRepository.FindByEmail(data.Email)
	if err != nil {
		WriteJSON(w, http.StatusNotFound, ResponseBody{
			Message: "account is not existed",
		})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(data.Password))
	if err != nil {
		WriteJSON(w, http.StatusUnauthorized, ResponseBody{
			Message: "email or password is incorrect",
			Code:    http.StatusUnauthorized,
		})
		return
	}
	claims := h.makeClaims((*account))
	token, err := h.CreateToken(claims)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ResponseBody{
			Message: "can not create token",
			Code:    http.StatusInternalServerError,
		})
		return
	}
	if err := h.AccessIDRepository.Set(account.Email, claims.Id, fmt.Sprintf("%d", claims.ExpiresAt)); err != nil {
		fmt.Println(err)
		WriteJSON(w, http.StatusInternalServerError, ResponseBody{
			Message: "Unable to create token",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	WriteJSON(w, http.StatusOK, ResponseBody{
		Message: "Successfully login",
		Data: map[string]interface{}{
			"token": token,
		},
	})
}

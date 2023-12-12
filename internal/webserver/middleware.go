package webserver

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/luckysetiawan/curio-qa-api/internal/constant"
	"github.com/luckysetiawan/curio-qa-api/internal/util"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

var jwtKey = []byte(os.Getenv("JWT_KEY"))
var tokenName = "token"

func GenerateToken(w http.ResponseWriter, r *http.Request, id string, username string, userType int) {
	tokenExpiryTime := time.Now().Add(constant.TokenExpiryTime * time.Minute)

	claims := &Claims{
		ID:       id,
		Username: username,
		UserType: userType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(tokenExpiryTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     tokenName,
		Value:    signedToken,
		Expires:  tokenExpiryTime,
		Secure:   false,
		HttpOnly: true,
	})
}

func ResetToken(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     tokenName,
		Value:    "",
		Expires:  time.Now(),
		Secure:   false,
		HttpOnly: true,
	})
}

func Authenticate(next http.HandlerFunc, accessType []int) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isValidToken := validateUserToken(r, accessType)
		if !isValidToken {
			NewJsonPresenter().SendUnathorized(w)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func validateUserToken(r *http.Request, accessType []int) bool {
	isAccessTokenValid, _, _, userType := validateTokenFromCookies(r)

	if isAccessTokenValid {
		isUserValid := util.Contains(accessType, userType)
		if isUserValid {
			return true
		}
	}
	return false
}

func validateTokenFromCookies(r *http.Request) (bool, string, string, int) {
	if cookie, err := r.Cookie(tokenName); err == nil {
		accessToken := cookie.Value
		accessClaims := &Claims{}
		parsedToken, err := jwt.ParseWithClaims(accessToken, accessClaims, func(accessToken *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err == nil && parsedToken.Valid {
			return true, accessClaims.ID, accessClaims.Username, constant.RegisteredUser
		}
	}
	return false, "", "", -1
}

// Package webserver provides the necessary functionality to run a server.
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

// init loads environment variables when the server starts.
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

// JWT key and token name.
var jwtKey = []byte(os.Getenv("JWT_KEY"))
var tokenName = "token"

// GenerateToken generates a new token and set the cookie.
func GenerateToken(w http.ResponseWriter, r *http.Request, id string, username string, userType int) {
	tokenExpirationTime := time.Now().Add(constant.TokenExpirationTime * time.Minute)

	claims := &Claims{
		ID:       id,
		Username: username,
		UserType: userType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(tokenExpirationTime),
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
		Expires:  tokenExpirationTime,
		Secure:   false,
		HttpOnly: true,
	})
}

// ResetToken clears the token from the cookie.
func ResetToken(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     tokenName,
		Value:    "",
		Expires:  time.Now(),
		Secure:   false,
		HttpOnly: true,
	})
}

// Authenticate checks whether the request is valid.
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

// validateUserToken checks whether the user type exists in allowed access type.
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

// validateTokenFromCookies gets the token then validate the token.
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

// GetDataFromCookies returns the ID and username data obtained from cookies.
func GetDataFromCookies(r *http.Request) (string, string) {
	if cookie, err := r.Cookie(tokenName); err == nil {
		accessToken := cookie.Value
		accessClaims := &Claims{}
		parsedToken, err := jwt.ParseWithClaims(accessToken, accessClaims, func(accessToken *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err == nil && parsedToken.Valid {
			return accessClaims.ID, accessClaims.Username
		}
	}
	return "", ""
}

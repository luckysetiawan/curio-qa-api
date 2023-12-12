package webserver

import (
	"github.com/golang-jwt/jwt/v5"
)

// JSON Web Token Claims
type Claims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	UserType int    `json:"user_type"`
	jwt.RegisteredClaims
}

// Response
type BaseResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Count   int         `json:"count,omitempty"`
}

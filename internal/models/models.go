package models

import (
	"net/http"
	"time"
)

type Middleware func(handler http.HandlerFunc) http.HandlerFunc

type Session struct {
	UserID         int       `json:"user_id"`
	ConnectionID   int       `json:"connection_id"`
	Username       string    `json:"username"`
	IpAddress      string    `json:"ip_address"`
	ExpirationTime time.Time `json:"expiration_time"`
}

type Credentials struct {
	Username string
	Password string
}

type User struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	HashedPwd string `json:"hash"`
	Salt      string `json:"salt"`
	Email     string `json:"email"`
}

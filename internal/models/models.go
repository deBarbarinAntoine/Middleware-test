package models

import (
	"net/http"
	"time"
)

type Middleware func(handler http.HandlerFunc) http.HandlerFunc

type Session struct {
	UserID         int
	ConnectionID   int
	Username       string
	IpAddress      string
	ExpirationTime time.Time
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

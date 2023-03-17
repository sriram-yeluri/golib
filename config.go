package golib

import "os"

type Auth struct {
	UserName string
	Password string
}

func SetAuth() *Auth {
	auth := new(Auth)
	auth.UserName = os.Getenv("AUTH_USERNAME")
	auth.Password = os.Getenv("AUTH_PASSWORD")
	return auth
}
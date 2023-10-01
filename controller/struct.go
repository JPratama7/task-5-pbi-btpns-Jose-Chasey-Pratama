package controller

import (
	"btpn/app"
	"btpn/middleware"
)

type Users struct {
	db  app.UserDatabase
	jwt *middleware.JWT
}

type Picture struct {
	db  app.PictureDatabase
	jwt *middleware.JWT
}

func NewUsers(db app.UserDatabase, jwt *middleware.JWT) *Users {
	return &Users{db, jwt}
}

func NewPicture(db app.PictureDatabase, jwt *middleware.JWT) *Picture {
	return &Picture{db, jwt}
}

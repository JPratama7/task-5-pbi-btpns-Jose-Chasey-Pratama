package app

import "btpn/models"

type UserDatabase interface {
	GetByUserPassword(string, string) (models.User, error)
	Update(data models.User) error
	Insert(data models.User) error
	Delete(data models.User) error
	GetById(string) (models.User, error)
}

type PictureDatabase interface {
	GetAll(string2 string) ([]models.Picture, error)
	GetById(string, string) (models.Picture, error)
	Update(models.Picture) error
	Insert(models.Picture) error
	Delete(string, string) error
}

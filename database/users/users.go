package users

import (
	"btpn/models"
	"context"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *Database {
	return &Database{db.Session(&gorm.Session{Context: context.Background()}).Model(&models.User{})}
}

func (d *Database) GetByUserPassword(user, pass string) (data models.User, err error) {
	err = d.db.
		WithContext(context.Background()).
		Where("email = ? AND password = ?", user, pass).
		Find(&data).
		Error
	return
}

func (d *Database) Update(data models.User) (err error) {
	err = d.db.
		WithContext(context.Background()).
		Where("ID = ?", data.ID).
		Updates(&data).
		Error
	return
}
func (d *Database) Insert(data models.User) (err error) {
	err = d.db.
		WithContext(context.Background()).
		Create(&data).
		Error
	return
}

func (d *Database) Delete(data models.User) (err error) {
	err = d.db.
		WithContext(context.Background()).
		Where("usernames = ? AND email = ? AND ID = ?", data.Usernames, data.Email, data.ID).
		Delete(&data).
		Error
	return

}

func (d Database) GetById(id string) (data models.User, err error) {
	err = d.db.
		WithContext(context.Background()).
		Where("ID = ?", id).
		First(&data).
		Error
	return

}

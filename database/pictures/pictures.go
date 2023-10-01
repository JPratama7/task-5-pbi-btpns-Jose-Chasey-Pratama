package pictures

import (
	"btpn/models"
	"context"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *Database {
	return &Database{db.Session(&gorm.Session{Context: context.Background()}).Model(&models.Picture{})}
}

func (d *Database) GetAll(uname string) (data []models.Picture, err error) {
	err = d.db.
		WithContext(context.Background()).
		Where("userId = ?", uname).
		Find(&data).
		Error
	return
}
func (d *Database) GetById(uname, id string) (data models.Picture, err error) {
	err = d.db.
		WithContext(context.Background()).
		First(&data).
		Where("userId = ?", uname).
		Where("ID = ?", id).
		Error
	return
}

func (d *Database) Update(data models.Picture) (err error) {
	err = d.db.
		WithContext(context.Background()).
		Updates(&data).
		Error
	return
}
func (d *Database) Insert(data models.Picture) (err error) {
	err = d.db.
		WithContext(context.Background()).
		Create(&data).
		Error
	return
}

func (d *Database) Delete(uname, id string) (err error) {
	err = d.db.
		WithContext(context.Background()).
		Where("ID = ?", id).
		Where("userId = ?", uname).
		Delete(&models.Picture{}).
		Error
	return
}

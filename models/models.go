package models

import (
	"time"
)

const TableNameUser = "users"
const TableNamePicture = "pictures"

// User mapped from table <users>
type User struct {
	ID        int32     `gorm:"column:ID;primaryKey;autoIncrement:true" json:"ID"`
	Usernames string    `gorm:"column:usernames;not null" json:"usernames"`
	Email     string    `gorm:"column:email;not null" json:"email"`
	Password  string    `gorm:"column:password;not null" json:"password"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:current_timestamp()" json:"created_at"`
	UpdateAt  time.Time `gorm:"column:update_at;not null" json:"update_at"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}

// Picture mapped from table <pictures>
type Picture struct {
	ID       int32  `gorm:"column:ID;primaryKey;autoIncrement:true" json:"ID"`
	Title    string `gorm:"column:title;not null" json:"title"`
	Caption  string `gorm:"column:caption;not null" json:"caption"`
	PhotoURL string `gorm:"column:photo_url;not null" json:"photo_url"`
	UserID   string `gorm:"column:userId;not null" json:"userId"`
}

// TableName Picture's table name
func (*Picture) TableName() string {
	return TableNamePicture
}

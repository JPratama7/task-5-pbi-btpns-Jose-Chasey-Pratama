package helpers

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

func CreateGormConnection(dbConfig string, minConn, maxConn int, timeout time.Duration) *gorm.DB {
	DB, err := gorm.Open(mysql.Open(dbConfig), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	DB.Statement.RaiseErrorOnNotFound = true

	if db, er := DB.DB(); er == nil {
		db.SetMaxIdleConns(minConn)
		db.SetMaxOpenConns(maxConn)
		db.SetConnMaxLifetime(timeout)
	}
	if err != nil {
		panic(err)
	}
	return DB
}

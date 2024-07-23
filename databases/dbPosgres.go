package databases

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBPostgres() *gorm.DB {
	dsn := "host=localhost user=postgres password= dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

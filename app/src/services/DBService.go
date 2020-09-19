package services

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/yoshi0202/grass-app/app/src/constant"
)

func ConnectGorm() *gorm.DB {
	c := constant.NewConst()
	dsn := fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local", c.DbUser, c.DbPassword, c.DbProtocol, c.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

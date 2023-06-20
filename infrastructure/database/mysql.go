package database

import (
	"fmt"
	"log"

	"github.com/saufiroja/online-course-api/config"
	"github.com/saufiroja/online-course-api/models/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql(conf *config.AppConfig) *gorm.DB {
	host := conf.Mysql.Host
	port := conf.Mysql.Port
	user := conf.Mysql.User
	pass := conf.Mysql.Pass
	name := conf.Mysql.Name

	//  parse time to gmt+0
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		pass,
		host,
		port,
		name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		log.Panic(err)
	}

	log.Println("success connect to mysql database!")

	return db
}

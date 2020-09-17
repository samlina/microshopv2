package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func CreateConnection() (*gorm.DB, error) {

	// 从系统环境变量获取数据库信息
	//host := os.Getenv("DB_HOST")
	//user := os.Getenv("DB_USER")
	//DBName := os.Getenv("DB_NAME")
	//password := os.Getenv("DB_PASSWORD")

	host := "localhost:3306"
	user := "test"
	DBName := "gxc_user"
	password := "test"

	return gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			user, password, host, DBName,
		),
	)
}

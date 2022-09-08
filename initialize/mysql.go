package initialize

import (
	"example.com/m/v2/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Mysql() *gorm.DB {
	m := global.CONFIG.Mysql
	if m.DBName == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN: m.Dsn(),
	}

	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
	if err != nil {
		return nil
	}
	return db
}

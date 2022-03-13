package initialize

import (
	"fmt"
	"github.com/akazwz/imgin/global"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitGormDB() {
	global.GORMDB = InitMySql()
	if global.GORMDB == nil {
		log.Fatalln("数据库初始化失败")
	}
}

// InitMySql 初始化 MySql
func InitMySql() *gorm.DB {
	dbMySqlUser := os.Getenv("DB_MYSQL_USER")
	dbMySqlPassword := os.Getenv("DB_MYSQL_PASSWORD")
	dbMySqlHost := os.Getenv("DB_MYSQL_HOST")
	dbMySqlName := os.Getenv("DB_MYSQL_NAME")

	/* 获取 dsn */
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=Local",
		dbMySqlUser,
		dbMySqlPassword,
		dbMySqlHost,
		dbMySqlName,
	)
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		return nil
	} else {
		//sqlDB, _ := db.DB()
		//sqlDB.SetMaxIdleConns()
		//sqlDB.SetMaxIdleConns()
		return db
	}
}

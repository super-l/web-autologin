package core

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

/*
ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION

# 修改全局
set @@global.sql_mode = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION';

# 修改当前
set @@sql_mode = 'STRICT_TRANS_TABLES,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION';

*/
func InitConnect() {
	var err error
	connstr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local", ConfigData.MysqlUser, ConfigData.MysqlPassword, ConfigData.MysqlHost, ConfigData.MysqlPort, ConfigData.MysqlDatabase)
	DB, err = gorm.Open("mysql", connstr)

	if err != nil {
		log.Fatal(err.Error())
	}

	// 全局禁用表名复数
	DB.SingularTable(true)
	DB.LogMode(false) // 开启SQL语句显示

	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
}

func CheckMysqlStatus() error {
	var err error
	connstr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local", ConfigData.MysqlUser, ConfigData.MysqlPassword, ConfigData.MysqlHost, ConfigData.MysqlPort, ConfigData.MysqlDatabase)
	db, err := gorm.Open("mysql", connstr)
	if err != nil {
		return err
	}
	db.Close()
	return nil
}

func Close() {
	var err error
	err = DB.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
}

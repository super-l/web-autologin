package core

import (
	"encoding/xml"
	"fmt"
	"os"
)

var ConfigData *Configuration

type Configuration struct {
	TaskHour   int    `xml:"task_hour"`
	ShowWindow int    `xml:"show_window"`
	UserName   string `xml:"username"`
	Password   string `xml:"password"`

	// mysql数据库
	MysqlUser     string `xml:"mysql_user"`
	MysqlPassword string `xml:"mysql_password"`
	MysqlHost     string `xml:"mysql_host"`
	MysqlPort     string `xml:"mysql_port"`
	MysqlDatabase string `xml:"mysql_database"`
}

func init() {
	// 读取配置文件
	configFile, errOpen := os.Open("config.xml")
	if errOpen != nil {
		fmt.Println("Error opening file:", errOpen)
		return
	}
	defer configFile.Close()
	if err := xml.NewDecoder(configFile).Decode(&ConfigData); err != nil {
		fmt.Println("Error Decode file:", err)
		return
	}
}

package main

import (
	"fmt"
	"svjia-cookie/core"
	"time"
)

func main() {
	fmt.Println("|------------------------------------------|")
	fmt.Println("|         三维家cookie自动获取               |")
	fmt.Println("|------------------------------------------|")
	fmt.Println("|      自动模拟登录，获取cookie信息           |")
	fmt.Println("|  Author: superl www.xiao6.net (小六博客)  |")
	fmt.Println("|------------------------------------------|")
	fmt.Println("")

	// 链接数据库
	core.InitConnect()
	defer core.Close()

	runService()
}

func runService() {
	// 先执行一次
	doService()

	var ticker = time.NewTicker(time.Duration(core.ConfigData.TaskHour) * time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			doService()
		}
	}
}

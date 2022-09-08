package main

import (
	"BlueBell/config"
	"BlueBell/db/mysql"
	"BlueBell/db/redis"
	"BlueBell/logger"
	"fmt"
)

func main() {
	// 1.加载配置文件
	if err := config.Init(); err != nil {
		fmt.Printf("config init failed!,err: %v\n", err)
		return
	}
	// 2.初始化日志文件
	if err := logger.Init(config.Conf.LogConfig, config.Conf.Mode); err != nil {
		fmt.Printf("config init failed!,err: %v\n", err)
		return
	}

	// 3.初始化MySQL
	if err := mysql.Init(config.Conf.MySQLConfig); err != nil {
		fmt.Printf("config init failed!,err: %v\n", err)
		return
	}
	defer mysql.Close()

	// 4.初始化Redis
	if err := redis.Init(config.Conf.RedisConfig); err != nil {
		fmt.Printf("config init failed!,err: %v\n", err)
		return
	}
	defer redis.Close()

	//5.优雅退出
}

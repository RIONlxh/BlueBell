package main

import (
	"BlueBell/config"
	"BlueBell/db/mysql"
	"BlueBell/db/redis"
	"BlueBell/logger"
	"BlueBell/routes"
	"BlueBell/utils/snowflake"
	"fmt"
)

func main() {
	// 1.加载配置文件
	if err := config.Init(); err != nil {
		fmt.Printf("init config failed!,err: %v\n", err)
		return
	}
	// 2.初始化日志文件
	if err := logger.Init(config.Conf.LogConfig, config.Conf.Mode); err != nil {
		fmt.Printf("init logger failed!,err: %v\n", err)
		return
	}

	// 3.初始化MySQL
	if err := mysql.Init(config.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed!,err: %v\n", err)
		return
	}
	defer mysql.Close()

	// 4.初始化Redis
	if err := redis.Init(config.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed!,err: %v\n", err)
		return
	}
	defer redis.Close()

	// 5.雪花算法
	if err := snowflake.Init(config.Conf.StartTime, config.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed!,err: %v\n", err)
		return
	}

	// 6.注册路由
	r := routes.SetUpRouter(config.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", config.Conf.Port))
	if err != nil {
		fmt.Printf("init routes failed!,err: %v\n", err)
		return
	}
}

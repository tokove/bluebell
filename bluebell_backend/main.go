package main

import (
	"bluebell_backend/dao/mysql"
	"bluebell_backend/dao/redis"
	"bluebell_backend/logger"
	"bluebell_backend/pkg/snowflake"
	"bluebell_backend/routers"
	"bluebell_backend/settings"
	"fmt"
)

func main() {
	//var confFile string
	//flag.StringVar(&confFile, "config", "./config/config.yaml", "配置文件")
	//flag.Parse()
	// viper配置
	if err := settings.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}

	// zap日志配置
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	// mysql配置
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close()

	// redis配置
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()

	// 雪花算法
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return 
	}

	// 注册路由
	r := routers.SetupRouter()
	err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}

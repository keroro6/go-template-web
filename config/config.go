package config

import "github.com/zeromicro/go-zero/rest"

//todo: 做自己的删改，如果不用的配置去掉
type Config struct {
	rest.RestConf
	Mysql struct {
		Dsn string
	}
	MongoDB []struct {
		DBName string
		Dsn    string
	}
}

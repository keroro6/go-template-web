package service

import (
	config2 "go-template-web/config"
)

type SrvContext struct {
	Config config2.Config
}

func NewSrvContext(c config2.Config) *SrvContext {
	//sqlx.NewMysql()
	//初始化mysql
	//if c.Mysql.DataSource != "" {
	//	sqlx.NewMysql(c.Mysql.DataSource)
	//}
	return &SrvContext{
		Config: c,
	}
}

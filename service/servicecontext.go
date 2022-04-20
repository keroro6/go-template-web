package service

import (
	config2 "go-template-web/config"
)

type SrvContext struct {
	Config config2.Config
}

func NewSrvContext(c config2.Config) *SrvContext {
	return &SrvContext{
		Config: c,
	}
}

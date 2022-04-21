package main

import (
	"flag"
	"fmt"
	config2 "go-template-web/config"
	"go-template-web/model/mongoDao"
	"go-template-web/service"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"go-template-web/controller"
)

var configFile = flag.String("f", "config/go-template-web-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config2.Config
	conf.MustLoad(*configFile, &c)
	ctx := service.NewSrvContext(c)
	initJob(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	controller.RegisterHandlers(server, ctx)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
	closeJob(c)
}

func initJob(c config2.Config) {
	fmt.Printf("%+v\n", c)
	mongoDao.InitMongo(c)
}

func closeJob(c config2.Config) {
	fmt.Println("closing...")
	mongoDao.CloseMongo()
}

package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	config2 "go-template-web/config"
	"go-template-web/controller"
	"go-template-web/controller/modelcontroller"
	"go-template-web/dao/mongoDao"
	"go-template-web/middleware"
	"go-template-web/service"
	"log"
	"net/http"
	"net/http/pprof"
)

var env = flag.String("e", "prod", "env")

const (
	EnvProd = "prod"
	EnvDev  = "dev"
)

func main() {
	flag.Parse()

	var c config2.Config

	configFile := "config/prod_config.yaml"
	if *env == EnvDev {
		configFile = "config/dev_config.yaml"
	}
	conf.MustLoad(configFile, &c)
	ctx := service.NewSrvContext(c)
	initJob(c)
	//跨域
	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()
	server.Use(middleware.TraceLog)
	controller.RegisterHandlers(server, ctx)
	//model路由
	modelcontroller.RegisterHandlers(server, ctx)
	//...
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
	closeJob(c)
}

func initJob(c config2.Config) {
	fmt.Printf("%+v\n", c)
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
		http.HandleFunc("/debug/pprof/", pprof.Index)
		http.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		http.HandleFunc("/debug/pprof/profile", pprof.Profile)
		http.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		http.HandleFunc("/debug/pprof/trace", pprof.Trace)
	}()
	mongoDao.InitMongo(c)
}

func closeJob(c config2.Config) {
	fmt.Println("closing...")
	mongoDao.CloseMongo()
}

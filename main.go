package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	config2 "go-template-web/config"
	"go-template-web/controller"
	"go-template-web/model/mongoDao"
	"go-template-web/service"
	"log"
	"net/http"
	"net/http/pprof"
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

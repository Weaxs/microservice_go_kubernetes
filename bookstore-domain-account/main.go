package main

import (
	"github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/api"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"log"
)

func main() {
	initLog()
	//addr, err := net.ResolveTCPAddr("tcp", ":8810")
	//if err != nil {
	//	panic(any(err))
	//}

	var options []server.Option
	// service address
	//options = append(options, server.WithServiceAddr(addr))
	// Multiplexing
	options = append(options, server.WithMuxTransport())

	svr := server.NewServer(options...)

	if err := svr.RegisterService(api.AccountApiServiceInfo, new(handler)); err != nil {
		panic(any(err))
	}

	if err := svr.Run(); err != nil {
		log.Println("server stopped with error: ", err)
	} else {
		log.Println("server stopped")
	}
}

func initLog() {
	klog.SetLogger(logrus.NewLogger())
	klog.SetLevel(klog.LevelDebug)
}

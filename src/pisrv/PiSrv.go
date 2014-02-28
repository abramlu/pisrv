package main

import (
	"base/log"
	"base/socket"
	"os"
)

var (
	Logger    *log.BeeLogger
	socketSrv *socket.Server
)

func main() {
	if err := initLogger(); err != nil {
		os.Exit(1)
	}

	if err := initSrv(); err != nil {
		os.Exit(1)
	}
}

//init global logger
func initLogger() error {
	Logger = log.NewLogger(1000)
	Logger.SetLogger(log.File, "{\"fileName\":\"pisrv.log\"}")
	Logger.Setlevel(log.Debug)
	return nil
}

//init socket server
func initSrv() error {
	addr := "0.0.0.0:10000"
	conf := socket.NewConfig()
	conf.Addr = addr
	conf.CodecFactory = socket.NewDefaultCodecFactory()

	conf.ConnectedHandler = connectedHandler
	conf.DisconnectHandler = disconnectHandler
	conf.MessageHandler = messageHandler

	var err error
	socketSrv, err = socket.NewServer(conf)
	if err != nil {
		return nil
	}
	socketSrv.Start()
	return nil

}

// 处理接受到的数据
func messageHandler(channel socket.IChannel, protoPack *socket.ProtoPack) {

}

//客户端连接事件处理器
func connectedHandler(channel socket.IChannel) {

}

//客户端连接断开事件处理器
func disconnectHandler(channel socket.IChannel) {

}

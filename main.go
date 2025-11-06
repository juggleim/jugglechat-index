package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/juggleim/jugglechat-index/apis"
	"github.com/juggleim/jugglechat-index/configures"
	"github.com/juggleim/jugglechat-index/dbs"
)

func main() {
	//init configure
	if err := configures.InitConfigures(); err != nil {
		fmt.Println("Init Configures failed", err)
		return
	}
	//init mysql
	if err := dbs.InitMysql(); err != nil {
		fmt.Println("Init Mysql failed", err)
		return
	}

	httpServer := gin.Default()
	httpServer.GET("/serverinfos", apis.GetServerInfo)

	go httpServer.Run(fmt.Sprintf(":%d", configures.Config.Port))

	closeChan := make(chan struct{})
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sigChan
		signal.Stop(sigChan)
		close(closeChan)
	}()

	<-closeChan
}

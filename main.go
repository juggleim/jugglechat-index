package main

import (
	"fmt"
	"net/http"
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
	httpServer.Use(corsHandler())
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

func corsHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		context.Writer.Header().Add("Access-Control-Allow-Headers", "*")
		context.Writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		context.Writer.Header().Add("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Writer.Header().Add("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}

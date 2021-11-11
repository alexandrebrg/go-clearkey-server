package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/protocole/clearkey/license"
	"gitlab.com/protocole/clearkey/logger"
	"go.uber.org/zap"
)

func RegisterLog() {
	zlogger, _ := zap.NewDevelopment()
	defer zlogger.Sync()

	sugar := zlogger.Sugar()
	logger.SetLogger(sugar)
}

func main() {
	RegisterLog()
	router := gin.Default()
	router.POST("/license", license.HandleRequest)
	router.POST("/license/register", license.HandleKeyRegistration)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}

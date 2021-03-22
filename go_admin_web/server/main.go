package main

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func main() {
	global.GVA_VP = initialize.Viper()
	global.GVA_LOG = initialize.Zap()
	global.GVA_DB = initialize.Gorm()
	if global.GVA_DB != nil {
		initialize.MysqlTables(global.GVA_DB)
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	initialize.Redis()
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	// gin.SetMode(gin.ReleaseMode)
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}

type server interface {
	ListenAndServe() error
}

func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

package initialize

import (
	"gin-vue-admin/global"
	"gin-vue-admin/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	// Router.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	// 跨域
	//Router.Use(middleware.Cors()) // 如需跨域可以打开
	//global.GVA_LOG.Info("use middleware cors")
	// 方便统一添加路由组前缀 多服务器上线使用
	//PublicGroup := Router.Group("")
	//{
	//	router.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	//	router.InitInitRouter(PublicGroup) // 自动初始化相关
	//}
	PrivateGroup := Router.Group("/wx")
	//PrivateGroup.Use(middleware.JWTAuth())
	{
		router.InitAuthRouter(PrivateGroup)
		router.InitUserRouter(PrivateGroup)
	}
	global.GVA_LOG.Info("router register success")
	return Router
}

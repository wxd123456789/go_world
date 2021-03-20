package router

import (
	v1 "gin-vue-admin/api/v1"
	"github.com/gin-gonic/gin"
)

func InitAuthRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("auth")
	{
		ApiRouter.GET("login", v1.Auth)
	}
}



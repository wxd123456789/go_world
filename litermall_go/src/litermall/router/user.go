package router

import (
	v1 "gin-vue-admin/api/v1"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("user")
	{
		ApiRouter.GET("", v1.GetAllUsers)
	}
}

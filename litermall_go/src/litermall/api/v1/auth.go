package v1

import (
	"gin-vue-admin/model/response"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	response.OkWithMessage("success", c)
}

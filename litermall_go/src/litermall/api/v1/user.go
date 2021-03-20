package v1

import (
	"fmt"
	"gin-vue-admin/dao"
	"gin-vue-admin/global"
	"gin-vue-admin/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetAllUsers(c *gin.Context) {
	err, userList := dao.GetAllUsers()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: userList,
		}, "获取成功", c)
	}
}

package dao

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)

func GetAllUsers() (err error, list interface{}) {
	db := global.GVA_DB.Model(&model.User{})
	var UserList []model.User
	err = db.Find(&UserList).Error
	if err != nil {
		return err, UserList
	}
	return err, UserList
}

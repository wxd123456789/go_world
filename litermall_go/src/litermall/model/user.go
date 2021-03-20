package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID            int           `gorm:"column:id;primary_key" json:"id"`               //
	Username      string        `gorm:"column:username" json:"username"`               //用户名称
	Password      string        `gorm:"column:password" json:"password"`               //用户密码
	Gender        int           `gorm:"column:gender" json:"gender"`                   //性别：0 未知， 1男， 1 女
	Birthday      time.Time     `gorm:"column:birthday" json:"birthday"`               //生日
	LastLoginTime time.Time     `gorm:"column:last_login_time" json:"last_login_time"` //最近一次登录时间
	LastLoginIP   string        `gorm:"column:last_login_ip" json:"last_login_ip"`     //最近一次登录IP地址
	UserLevel     sql.NullInt64 `gorm:"column:user_level" json:"user_level"`           //0 普通用户，1 VIP用户，2 高级VIP用户
	Nickname      string        `gorm:"column:nickname" json:"nickname"`               //用户昵称或网络名称
	Mobile        string        `gorm:"column:mobile" json:"mobile"`                   //用户手机号码
	Avatar        string        `gorm:"column:avatar" json:"avatar"`                   //用户头像图片
	WeixinOpenid  string        `gorm:"column:weixin_openid" json:"weixin_openid"`     //微信登录openid
	SessionKey    string        `gorm:"column:session_key" json:"session_key"`         //微信登录会话KEY
	Status        int           `gorm:"column:status" json:"status"`                   //0 可用, 1 禁用, 2 注销
	AddTime       time.Time     `gorm:"column:add_time" json:"add_time"`               //创建时间
	UpdateTime    time.Time     `gorm:"column:update_time" json:"update_time"`         //更新时间
	Deleted       sql.NullInt64 `gorm:"column:deleted" json:"deleted"`                 //逻辑删除
}

// TableName sets the insert table name for this struct type
func (t *User) TableName() string {
	return "litemall_user"
}

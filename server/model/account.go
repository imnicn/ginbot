/*
 * @Author: yesxin(陶鑫)
 * @Github: https://github.com/yesxin
 * @Date: 2020-11-10 19:28:32
 */

package model

import (
	"time"
)

/*Account 用户表*/
type Account struct {
	ID        uint32    `gorm:"comment:'账户ID'"`
	Mobil     string    `gorm:"size:16;comment:'手机'"`
	Email     string    `gorm:"size:32;comment:'邮箱'"`
	Passwd    string    `gorm:"size:32;comment:'密码'"`
	Name      string    `gorm:"size:64;comment:'名称'"`
	Gender    uint8     `gorm:"comment:'性别 0：未知、1：男、2：女'"`
	Birthday  time.Time `gorm:"comment:'生日'"`
	Profile   string    `gorm:"size:255;comment:'个人简介'"`
	Avatar    string    `gorm:"size:255;comment:'头像'"`
	Enable    bool      `gorm:"comment:'启用状态'"`
	CreatedAt time.Time `gorm:"comment:'创建时间'"`
	UpdatedAt time.Time `gorm:"comment:'更新时间'"`
	Bots      []Bot
}

//Login 小程序登录
func (m *Account) Login() error {
	if err := db.Where(m).Take(m).Error; err != nil {
		return err
	}
	return nil
}

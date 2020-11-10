/*
 * @Author: yesxin(陶鑫)
 * @Github: https://github.com/yesxin
 * @Date: 2020-11-09 23:17:01
 */

package model

//Bot 微信机器人表
type Bot struct {
	ID        uint32 `gorm:"comment:'机器人ID'"`
	Token     string `gorm:"size:32;comment:'机器人Token'"`
	Endpoint  string `gorm:"size:32;comment:'机器人Endpoint'"`
	AccountID uint32 `gorm:"comment:'所属账户ID'"`
}

//Get 根据主键获取数据
func (m *Bot) Get() error {
	return db.Take(m).Error
}

/*
 * @Author: yesxin(陶鑫)
 * @Github: https://github.com/yesxin
 * @Date: 2020-11-09 20:32:25
 */

package model

import (
	"fmt"
	"ginbot/app"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func init() {
	log.Println("数据库开始连接……")
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		app.Config.MysqlUser,
		app.Config.MysqlPasswd,
		app.Config.MysqlHost,
		app.Config.MysqlPort,
		app.Config.MysqlDbname,
		app.Config.MysqlCharset)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		panic("数据库连接错误")
	}
	log.Println("数据库连接成功。")

	db.AutoMigrate(&Account{}, &Bot{})
}

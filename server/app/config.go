/*
 * @Author: yesxin(陶鑫)
 * @Github: https://github.com/yesxin
 * @Date: 2020-11-08 20:46:48
 */

package app

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

//config 配置
type config struct {
	MysqlUser    string `json:"mysql_user"`
	MysqlPasswd  string `json:"mysql_passwd"`
	MysqlHost    string `json:"mysql_host"`
	MysqlPort    string `json:"mysql_port"`
	MysqlDbname  string `json:"mysql_dbname"`
	MysqlCharset string `json:"mysql_charset"`
}

//Config 配置实例
var Config = new(config)

//ReloadConfig 重载配置
func ReloadConfig() error {
	if err := readFileConfig(); err != nil {
		return err
	}
	if err := readDbConfig(); err != nil {
		return err
	}
	return nil
}

func init() {
	log.Println("配置文件开始加载……")
	if err := readFileConfig(); err != nil {
		panic(err)
	}
	if err := readDbConfig(); err != nil {
		panic(err)
	}
	log.Println("配置文件加载完成")
}

func readFileConfig() error {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		return errors.New("配置文件读取错误，找不到配置文件")
	}

	if err = json.Unmarshal(file, Config); err != nil {
		return errors.New("配置文件解析失败" + err.Error())
	}
	return nil
}

func readDbConfig() error {
	return nil
}

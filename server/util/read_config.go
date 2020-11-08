/*
 * @Author: yesxin(陶鑫)
 * @Github: https://github.com/yesxin
 * @Date: 2020-11-08 20:46:48
 */

package util

import (
	"encoding/json"
	"io/ioutil"
)

//Config 配置
type Config struct {
	Database string `json:"database"`
}

//Conf 配置
var Conf Config

func init() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic("配置文件读取错误，找不到配置文件")
	}

	if err = json.Unmarshal(file, &Conf); err != nil {
		panic("配置文件解析失败")
	}
}

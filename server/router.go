/*
 * @Author: yesxin(陶鑫)
 * @Github: https://github.com/yesxin
 * @Date: 2020-11-08 19:40:22
 */

package main

import (
	"ginbot/handler/bot"
	"ginbot/handler/home"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	groupHome := r.Group("/home")
	groupHome.GET("/", home.HandlerGetIndex)

	groupBot := r.Group("/bot")
	groupBot.GET("/start", bot.Start)
	groupBot.GET("/stop", bot.Stop)
	groupBot.GET("/find_all_room", bot.FindAllRoom)

	return r
}

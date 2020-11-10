/*
 * @Author: yesxin(陶鑫)
 * @Github: https://github.com/yesxin
 * @Date: 2020-11-08 19:40:22
 */

package router

import (
	"ginbot/handler/auth"
	"ginbot/handler/bot"
	"ginbot/handler/home"
	"ginbot/midware"

	"github.com/gin-gonic/gin"
)

//SetupRouter SetupRouter
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/login", auth.Login)

	groupHome := r.Group("/home")
	groupHome.GET("/", home.HandlerGetIndex)

	groupBot := r.Group("/bot", midware.AuthJWT())
	groupBot.GET("/start", bot.Start)
	groupBot.GET("/stop", bot.Stop)
	groupBot.GET("/logout", bot.Logout)
	groupBot.GET("/find_all_room", bot.FindAllRoom)

	return r
}

/*
 * @Author: yesxin(陶鑫)
 * @Github: https://github.com/yesxin
 * @Date: 2020-11-08 19:59:11
 */

package bot

import (
	"ginbot/app"
	"ginbot/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*Start 机器人启动*/
func Start(c *gin.Context) {
	bot := new(model.Bot)
	bot.ID = uint(1)
	if err := bot.Get(); err != nil {
		c.JSON(http.StatusForbidden, err)
	}
	if err := app.Bot.Start(bot.Token, bot.Endpoint); err != nil {
		c.JSON(http.StatusForbidden, err)
	}

	c.JSON(http.StatusOK, app.Bot.GetID(bot.Token))
}

/*Stop 机器人停止*/
func Stop(c *gin.Context) {
	bot := new(model.Bot)
	bot.ID = uint(1)
	if err := bot.Get(); err != nil {
		c.JSON(http.StatusForbidden, err)
	}

	err := app.Bot.Stop(bot.Token)

	c.JSON(http.StatusOK, err.Error()+app.Bot.GetID(bot.Token))
}

/*Logout 机器人注销*/
func Logout(c *gin.Context) {
	bot := new(model.Bot)
	bot.ID = uint(1)
	if err := bot.Get(); err != nil {
		c.JSON(http.StatusForbidden, err)
	}

	err := app.Bot.Logout(bot.Token)

	c.JSON(http.StatusOK, err.Error()+app.Bot.GetID(bot.Token))
}

/*FindAllRoom 查找全部微信群*/
func FindAllRoom(c *gin.Context) {
	bot := new(model.Bot)
	bot.ID = uint(1)
	if err := bot.Get(); err != nil {
		c.JSON(http.StatusForbidden, err)
	}

	c.JSON(http.StatusOK, app.Bot.FindAllRoom(bot.Token))
}

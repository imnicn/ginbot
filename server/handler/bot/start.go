/*
 * @Author: yesxin(陶鑫)
 * @Github: https://github.com/yesxin
 * @Date: 2020-11-08 19:59:11
 */

package bot

import (
	"ginbot/extend/wechaty"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*Start 机器人启动*/
func Start(c *gin.Context) {
	if err := wechaty.BotStart("11111111", "ip:port"); err != nil {
		c.JSON(http.StatusForbidden, err)
	}

	c.JSON(http.StatusOK, nil)
}

/*Stop 机器人停止*/
func Stop(c *gin.Context) {
	err := wechaty.BotStop("11111111")

	c.JSON(http.StatusOK, err)
}

/*FindAllRoom 查找全部微信群*/
func FindAllRoom(c *gin.Context) {
	c.JSON(http.StatusOK, wechaty.FindAllRoom("11111111"))
}

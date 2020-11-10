/*
 * @Author: yesxin(陶鑫)
 * @Github: https://github.com/yesxin
 * @Date: 2020-11-10 19:28:52
 */

package midware

import (
	"ginbot/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

//AuthJWT 验证JWT
func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := c.Request.Header.Get("Authorization")
		claims, err := app.ParseJWT(t)

		if err != nil {
			c.JSON(http.StatusUnauthorized, "登录状态验证失败，请先登录。")
			c.Abort()
			return
		}
		c.Set("AccountID", claims.Id)
		c.Next()
	}
}

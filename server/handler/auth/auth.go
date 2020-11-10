/*
 * @Author: yesxin(陶鑫)
 * @Github: https://github.com/yesxin
 * @Date: 2020-11-10 20:38:19
 */

package auth

import (
	"ginbot/app"
	"ginbot/model"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

/*Login Login*/
func Login(c *gin.Context) {
	body := new(struct {
		Acc string `json:"acc" binding:"required"`
		Pwd string `json:"pwd" binding:"required"`
	})
	if err := c.ShouldBind(body); err != nil {
		c.JSON(http.StatusBadRequest, "请求Json绑定失败。")
		return
	}

	isMobile, _ := regexp.MatchString("^(13[0-9]|14[5|7]|15[0-9]|17[0-9]|18[0-9]|19[0-9])\\d{8}$", body.Acc)
	isEmail, _ := regexp.MatchString("^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$", body.Acc)
	account := new(model.Account)
	if isMobile {
		account.Mobil = body.Acc
	} else if isEmail {
		account.Email = body.Acc
	} else {
		c.JSON(http.StatusBadRequest, "输入账号的格式错误。")
		return
	}
	account.Passwd = body.Pwd

	if err := account.Login(); err != nil {
		c.JSON(http.StatusUnauthorized, "输入的账户或密码错误。")
		return
	}
	token, err := app.CreateJWT(account.ID, "sub", time.Now().Unix()+24*60*60)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Token生成失败。")
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

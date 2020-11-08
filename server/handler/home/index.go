/*
 * @Author: yesxin(陶鑫)
 * @Github: https://github.com/yesxin
 * @Date: 2020-11-08 19:59:11
 */

package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*HandlerGetIndex HandlerGetIndex*/
func HandlerGetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, "hello")
}

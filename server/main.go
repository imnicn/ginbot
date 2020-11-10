/*
 * @Author: yesxin(陶鑫)
 * @Github: https://github.com/yesxin
 * @Date: 2020-11-08 19:17:27
 */

package main

import (
	"ginbot/app"
	"ginbot/router"
)

func main() {
	app.Start()

	gin := router.SetupRouter()
	gin.Run(":8080")
}

/*
 * @Author: yesxin(陶鑫)
 * @Github: https://github.com/yesxin
 * @Date: 2020-11-08 21:16:08
 */

package wechaty

import (
	"fmt"

	"github.com/wechaty/go-wechaty/wechaty"
	wp "github.com/wechaty/go-wechaty/wechaty-puppet"
	"github.com/wechaty/go-wechaty/wechaty-puppet/schemas"
	"github.com/wechaty/go-wechaty/wechaty/user"
)

//BotMap 机器人map
var BotMap = make(map[string]*wechaty.Wechaty)

//BotStart 启动机器人
func BotStart(token string, endpoint string) error {

	BotMap[token] = wechaty.NewWechaty(wechaty.WithPuppetOption(wp.Option{
		Endpoint: endpoint,
		Token:    token,
	}))

	BotMap[token].OnScan(func(ctx *wechaty.Context, qrCode string, status schemas.ScanStatus, data string) {
		//fmt.Println(ctx)
		fmt.Printf("https://wechaty.github.io/qrcode/%s\n", qrCode)
		fmt.Println(status)
		//fmt.Println(data)
	})

	BotMap[token].OnReady(func(ctx *wechaty.Context) {
		fmt.Println("OnReady")
	})

	BotMap[token].OnStop(func(ctx *wechaty.Context) {
		fmt.Println("onstop")
	})

	BotMap[token].OnError(func(ctx *wechaty.Context, err error) {
		fmt.Println(err)
	})

	BotMap[token].OnMessage(func(ctx *wechaty.Context, msg *user.Message) {
		//fmt.Println(msg)
	})

	if err := BotMap[token].Start(); err != nil {
		return err
	}
	return nil
}

func FindAllRoom(token string) []string {
	filter := new(schemas.RoomQueryFilter)
	rooms := BotMap[token].Room().FindAll(filter)

	var r []string
	for _, room := range rooms {
		r = append(r, room.Topic())
	}
	return r
}

func BotStop(token string) error {
	//BotMap[token].Puppet().Logout()
	BotMap[token].Puppet().Stop()
	return nil
}

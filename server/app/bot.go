/*
 * @Author: yesxin(陶鑫)
 * @Github: https://github.com/yesxin
 * @Date: 2020-11-08 21:16:08
 */

package app

import (
	"fmt"

	"github.com/wechaty/go-wechaty/wechaty"
	wp "github.com/wechaty/go-wechaty/wechaty-puppet"
	"github.com/wechaty/go-wechaty/wechaty-puppet/schemas"
	"github.com/wechaty/go-wechaty/wechaty/user"
)

type bot map[string]*wechaty.Wechaty

//Bot 机器人MAP实例
var Bot = make(bot)

//Start 启动机器人
func (b bot) Start(token string, endpoint string) error {

	b[token] = wechaty.NewWechaty(wechaty.WithPuppetOption(wp.Option{
		Endpoint: endpoint,
		Token:    token,
	}))

	b[token].OnScan(func(ctx *wechaty.Context, qrCode string, status schemas.ScanStatus, data string) {
		//fmt.Println(ctx)
		fmt.Printf("https://wechaty.github.io/qrcode/%s\n", qrCode)
		fmt.Println(status)
		//fmt.Println(data)
	})

	b[token].OnReady(func(ctx *wechaty.Context) {
		fmt.Println("OnReady")
	})

	b[token].OnStop(func(ctx *wechaty.Context) {
		fmt.Println("onstop")
	})

	b[token].OnError(func(ctx *wechaty.Context, err error) {
		fmt.Println(err)
	})

	b[token].OnMessage(func(ctx *wechaty.Context, msg *user.Message) {
		//fmt.Println(msg)
	})

	if err := b[token].Start(); err != nil {
		return err
	}
	return nil
}

//FindAllRoom 查询全部微信群
func (b bot) FindAllRoom(token string) []string {
	filter := new(schemas.RoomQueryFilter)
	rooms := b[token].Room().FindAll(filter)

	var r []string
	for _, room := range rooms {
		r = append(r, room.Topic())
	}
	return r
}

//Stop 停止机器人
func (b bot) Stop(token string) error {
	b[token].Puppet().Stop()
	return nil
}

//Logout 注销机器人
func (b bot) Logout(token string) error {
	return b[token].Puppet().Logout()
}

//GetID 获取机器人ID
func (b bot) GetID(token string) string {
	return b[token].Puppet().SelfID()
}

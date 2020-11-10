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

type bot map[uint32]*wechaty.Wechaty

//Bot 机器人MAP实例
var Bot = make(bot)
var qrcode = make(map[uint32]chan string)

//Start 启动机器人
func (b bot) Start(id uint32, token string, endpoint string) error {

	b[id] = wechaty.NewWechaty(wechaty.WithPuppetOption(wp.Option{
		Endpoint: endpoint,
		Token:    token,
	}))

	b[id].OnScan(func(ctx *wechaty.Context, qrCode string, status schemas.ScanStatus, data string) {
		if qrCode != "" {
			if qrcode[id] == nil {
				qrcode[id] = make(chan string)
			}
			qrcode[id] <- qrCode
		}
	})

	b[id].OnReady(func(ctx *wechaty.Context) {
		fmt.Println("OnReady")
	})

	b[id].OnStop(func(ctx *wechaty.Context) {
		fmt.Println("onstop")
	})

	b[id].OnError(func(ctx *wechaty.Context, err error) {
		fmt.Println("onerr")
	})

	b[id].OnMessage(func(ctx *wechaty.Context, msg *user.Message) {
		//fmt.Println(msg)
	})

	if err := b[id].Start(); err != nil {
		return err
	}
	return nil
}

//FindAllRoom 查询全部微信群
func (b bot) FindAllRoom(id uint32) []string {
	filter := new(schemas.RoomQueryFilter)
	rooms := b[id].Room().FindAll(filter)

	var r []string
	for _, room := range rooms {
		r = append(r, room.Topic())
	}
	return r
}

//Stop 停止机器人
func (b bot) Stop(id uint32) error {
	b[id].Puppet().Stop()
	return nil
}

//Logout 注销机器人
func (b bot) Logout(id uint32) error {
	return b[id].Puppet().Logout()
}

//GetID 获取机器人ID
func (b bot) GetID(id uint32) string {
	return b[id].Puppet().SelfID()
}

//GetQrcode 获取机器人二维码
func (b bot) GetQrcode(id uint32) string {
	if qrcode[id] == nil {
		qrcode[id] = make(chan string)
	}
	b[id].UserSelf().Weixin()
	return <-qrcode[id]
}

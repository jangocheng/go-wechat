package server

import (
	"github.com/MwlLj/go-wechat/common"
)

type CMsgCallbackDefault struct {
	MsgCallback common.FuncMsgCallback
}

func (this *CMsgCallbackDefault) OnMessage(reply common.IReply, msg *common.CMessage, userData interface{}) error {
	return this.MsgCallback(reply, msg, userData)
}

type CEventCallbackDefault struct {
	EventCallback         common.IEvent
	EventCallbackUserdata interface{}
}
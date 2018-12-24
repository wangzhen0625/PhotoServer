package main

import (
	"encoding/json"
	"time"

	"github.com/nsqio/go-nsq"
)

/*
filename  filebase64
*/
func StartReceiver() {
	ih := &ImgHandler{}
	ih.InitResources()
}

type ImgHandler struct {
}

type BodyMsg struct {
	TargetUrl  string // ip/a/b/c.jpg
	Filebase64 string //img
}

func (ih *ImgHandler) InitResources() {
	if imgConsumer, err := nsq.NewConsumer("img_server", "imgsave", nsq.NewConfig()); err != nil {
		SysLogger.Error("New imgConsumer err:" + err.Error())
		return
	} else {
		// 连接nsq
		imgConsumer.AddHandler(ih)                                                         // 设置消息处理函数
		if err := imgConsumer.ConnectToNSQD(Config.Services.Web.Nsqdtcpaddr); err != nil { // 连接到单例nsqd
			SysLogger.Error("Connect imgConsumer err:" + err.Error())
			return
		}
	}
}

// 这是自动调用
func (ih *ImgHandler) HandleMessage(msg *nsq.Message) error {
	//////////获取到log//////////
	defer TimeCost("---HandleMessage本次执行时间：", time.Now())
	bmsg := &BodyMsg{}
	if err := json.Unmarshal(msg.Body, bmsg); err != nil {
		SysLogger.Error("HandleMessage Unmarshal err:" + err.Error())
		return nil
	} else {
		postFile(bmsg.TargetUrl, bmsg.Filebase64)
		// bmsg.Filepath
		// bmsg.Filebase64
	}
	return nil
}

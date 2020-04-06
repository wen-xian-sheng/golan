package handle

import (
	"go_code/chatroom/common/message"
	"encoding/json"
	"fmt"
)

type SmsMgr struct{}

var smsMgr *SmsMgr

func(this *SmsMgr) printSendGroupMsg(msg *message.Message){
	//取出UserId,UserName和Content
	data := msg.Data
	var sendGroupMsg message.SendGroupMsg
	err := json.Unmarshal([]byte(data),&sendGroupMsg)
	if err != nil{
		fmt.Println("err := json.Unmarshal([]byte(data),&sendGroupMsg) err =",err)
		return
	}
	//输出UserId,UserName,Content
	fmt.Printf("接收到的群发消息：userId = %d,userName = %s,Content = %s",sendGroupMsg.UserId,sendGroupMsg.UserName,sendGroupMsg.Content)
}
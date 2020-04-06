package handle

import (
	"go_code/chatroom/common/message"
	"go_code/chatroom/utils"
	"encoding/json"
)

type SmsHandle struct{

}

func(this *SmsHandle) SendGroupMsg(msg *message.Message)(err error){
	//取出发起群聊的userId,不通知自己发送群聊的提示
	data := msg.Data
	var sendGroupMsg message.SendGroupMsg
	err = json.Unmarshal([]byte(data),&sendGroupMsg)
	curUserId := sendGroupMsg.UserId
	delete(userMgr.onlineUser,curUserId)

	//通知每个在线用户有一个用户发送群聊消息
	var bytes []byte 
	bytes,err = json.Marshal(msg)
	if err != nil{
		return
	}

	//遍历在线用户
	for _,v := range userMgr.onlineUser{
		
		//响应给客户端
		transfer := utils.Transfer{
			Conn : v.Conn,
		}
		err = transfer.WritePkg(bytes)
		}

	return
}
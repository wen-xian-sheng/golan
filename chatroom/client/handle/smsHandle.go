package handle
import (
	"go_code/chatroom/common/message"
	"go_code/chatroom/utils"
	"fmt"
)

type SmsHandle struct{

}

//一个客户端只有一个实例
var smsHandle SmsHandle

//群发消息
func(this *SmsHandle) sendGroupMsg(content string)(err error){
	fmt.Println("开始发送消息")
	//构建SendGroupMsg
	var sendGroupMsg message.SendGroupMsg
	sendGroupMsg.UserId = curUser.UserId
	sendGroupMsg.UserState = curUser.UserState
	sendGroupMsg.Content = content

	//向服务器发送package
	transfer := &utils.Transfer{
		Conn : curUser.Conn,
	}
	err = transfer.MarshalAndWritePkg(sendGroupMsg,message.SendGroupMsgType)
	if err != nil{
		fmt.Println("群发消息时，向服务器writePkg失败 err =",err)
		return
	}
	return
}
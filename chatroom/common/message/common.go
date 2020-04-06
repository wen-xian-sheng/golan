package message
import (
	"go_code/chatroom/common/model"
)

//消息type
const (
	LoginMsgType         = "LoginMsg"
	LoginResMsgType      = "LoginResMsg"
	RegistMsgType        = "RegistMsg"
	RegistResMsgType     = "RegistResMsg"
	NotifyStateMsgType   = "NotifyStateMsg"
	SendGroupMsgType     = "SendGroupMsg"
)

//用户状态
const (
	USER_ONLINE = iota
	USER_BUSY
	USER_OFFLINE
)

type Message struct{
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMsg struct{
	UserId int `json:"userId"`
	UserPwd string  `json:"userPwd"`
	UserName string `json:"userName"`
}

type RegistMsg struct{
	User model.User `json:"user"`
}

type NotifyStateMsg struct{
	UserId int `json:"userId"`
	State int `json:"state"`
}
//群聊
type SendGroupMsg struct{
	Content string `json:"content"`
	model.User
}

type LoginResMsg struct{
	Code int `json:"code"`
	Err string `json:"err"`
	OnlineUserIds []int `json:"onlineUserIds"`
}

type RegistResMsg struct{
	Code int `json:"code"`
	Err string `json:"err"`
}
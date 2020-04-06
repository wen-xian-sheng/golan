package handle
import (
	"net"
	"go_code/chatroom/utils"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/model"
	"encoding/json"
	"fmt"
)

type UserHandle struct{
	Conn net.Conn
	UserId int
}


//用户注册
func (this *UserHandle) Regist(msg *message.Message)( err error){
	//反序列出message.RegistMsg
	var registMsg message.RegistMsg
	err = json.Unmarshal([]byte(msg.Data),&registMsg)
	if err != nil{
		return
	}

	//取出registMsg.User,注册到redis
	err = model.MyUserDao.Insert(&registMsg.User)

	//根据user注册结果 构建响应信息RegistRes
	var registResMsg message.RegistResMsg
	if err != nil{
		switch err {
		case model.ERROR_USER_EXIT:
			registResMsg.Code = 500
			registResMsg.Err = err.Error()
		default:
			registResMsg.Code = 505
			registResMsg.Err = "服务器内部错误"
		}
	}else{
		registResMsg.Code = 200
		fmt.Println("注册成功")
	}
	//发送包给客户端
	transfer := utils.Transfer{
		Conn : this.Conn,
	}
	err = transfer.MarshalAndWritePkg(registResMsg,message.RegistResMsgType)
	return
}

//验证登录并响应
func(this *UserHandle) Login(msg *message.Message)(err error){
	//反序列出message.LoginMsg
	var loginMsg message.LoginMsg
	err = json.Unmarshal([]byte(msg.Data),&loginMsg)
	if err != nil{
		return
	}
	//根据user验证结果 构建响应信息LoginResMsg
	var loginResMsg message.LoginResMsg
	//验证用户是否正确
	var user *model.User
	user,err = model.MyUserDao.GetByUserId(loginMsg.UserId,loginMsg.UserPwd) 
	if err != nil{
		switch err {
		case model.ERROR_USER_NOT_EXIT:
			loginResMsg.Code = 500
			loginResMsg.Err = err.Error()
		case model.ERROR_USERPWD:
			loginResMsg.Code = 403
			loginResMsg.Err = err.Error()
		default:
			loginResMsg.Code = 505
			loginResMsg.Err = "服务器内部错误"
		}
	}else{
		loginResMsg.Code = 200
		fmt.Println(user,"登录成功")

		//登录成功,将该UserHandle加入到在线用户map userMgr.onlineUser
		//标记该userHandle属于哪个用户Id
		this.UserId = loginMsg.UserId
		//将该UserHandle加入到在线用户map userMgr.onlineUser
		userMgr.insertOnlineUser(this)
		
		//构建该用户状态信息
		var notifyStateMsg message.NotifyStateMsg
		notifyStateMsg.UserId = loginMsg.UserId
		notifyStateMsg.State = message.USER_ONLINE
		//获取全部在线用户
		onlineUserMaps := userMgr.getAllOnlineUser()
		// 遍历onlineUserMaps
		for k,v := range onlineUserMaps{

			//如果在线用户是自己，就不用通知自己已经上线
			if k == loginMsg.UserId{
				continue
			}

		// //登录成功，获取全部在线用户的userId
			//获取在线用户的userId(即onlineUserMaps的key),并append到loginResMsg.OnlineUserIds
			loginResMsg.OnlineUserIds = append(loginResMsg.OnlineUserIds,k)

		//登录成功
			// 开一个goloutine通知在线用户该用户上线了
			go this.notifyOtherOnlineUserThisUserState(&notifyStateMsg,v)
		}
	}

	//发送package给客户端
	transfer := utils.Transfer{
		Conn : this.Conn,
	}
	err = transfer.MarshalAndWritePkg(loginResMsg,message.LoginResMsgType)
	return
}

// 通知一个在线用户该用户上线了
func(this *UserHandle) notifyOtherOnlineUserThisUserState(notifyStateMsg *message.NotifyStateMsg,
	userHandle *UserHandle){
		
		transfer := utils.Transfer{
			Conn : userHandle.Conn,
		}
		
		err := transfer.MarshalAndWritePkg(notifyStateMsg,message.NotifyStateMsgType)
		if err != nil{
			fmt.Println("userHandle.go notifyOtherOnlineUserThisUserState() err =",err)
		}
}

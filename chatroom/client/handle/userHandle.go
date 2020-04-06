package handle
import (
	"fmt"
	"net"
	"encoding/json"
	"go_code/chatroom/common/message"
	"go_code/chatroom/client/model"
	"go_code/chatroom/utils"
	"errors"
)

type UserHandle struct{
	
}
///一个客户端只有一个当前用户
var curUser model.CurUser

//获取user输入的userId userPwd
func(this *UserHandle) getScan() (userId int,userPwd string){
	fmt.Print("请输入userId: ")
	fmt.Scanln(&userId)
	fmt.Print("请输入userPwd: ")
	fmt.Scanln(&userPwd)
	return 
}

//获取用户输入的userId userPwd userName,并注册
func(this *UserHandle) Regist() (err error){
	
	var (
		userId int
		userPwd string
		userName string
	)
	userId,userPwd = this.getScan()
	fmt.Print("请输入userName: ")
	fmt.Scanln(&userName)

	//封装注册信息
	//构建RegistMsg
	var registMsg message.RegistMsg
	registMsg.User.UserId = userId
	registMsg.User.UserPwd = userPwd
	registMsg.User.UserName = userName
	//marshal registMsg
	var data []byte
	data,err = json.Marshal(registMsg)
	if err != nil{
		return
	}
	//构建Message
	var msg message.Message
	msg.Type = message.RegistMsgType
	msg.Data = string(data)
	//marshal Message
	data,err = json.Marshal(msg)
	if err != nil{
		return
	}

	//连接服务器
	var conn net.Conn
	conn,err = net.Dial("tcp","localhost:8889")
	defer conn.Close()
	//连接服务器失败，return
	if err != nil{
		return
	}
	//向服务器发送注册信息，并接受返回的信息
	msg,err = sendPkgAndRecievePkg(&conn,data)
	if err != nil{
		return
	}

	//unmarshal(msg) 反序列化成 message.RegistResMsg
	var registResMsg message.RegistResMsg
	err = json.Unmarshal([]byte(msg.Data),&registResMsg)
	if err != nil{
		return
	}
	if registResMsg.Code != 200{
		err = errors.New(registResMsg.Err)
		return
	}

	return
}

//登录
func(this *UserHandle) Login() (err error){
	userId,userPwd := this.getScan()
	err = this.VertifyUser(userId,userPwd)
	return
}

//向服务器验证user是否正确
func(this *UserHandle) VertifyUser(userId int,userPwd string) (err error) {

	//构建LoginMsg
	var loginMsg message.LoginMsg
	loginMsg.UserId = userId
	loginMsg.UserPwd = userPwd
	//marshal loginMsg
	var data []byte
	data,err = json.Marshal(loginMsg)
	if err != nil{
		return
	}
	//构建Message
	var msg message.Message
	msg.Type = message.LoginMsgType
	msg.Data = string(data)
	//marshal Message
	data,err = json.Marshal(msg)
	if err != nil{
		return
	}
	
	//连接服务器
	var conn net.Conn
	conn,err = net.Dial("tcp","localhost:8889")
	defer conn.Close()
	//连接服务器失败，return
	if err != nil{
		return
	}
	//向服务器发送并接收package
	msg,err = sendPkgAndRecievePkg(&conn,data)
	if err != nil{
		fmt.Println("userHandle.go VertifyUser() sendPkgAndRecievePkg(data)  err =",err)
		return
	}
	fmt.Println("接收到服务器返回的json字符串 msg =",msg)

	//unmarshal(msg) 反序列化成 message.LoginResMsg
	var loginResMsg message.LoginResMsg
	err = json.Unmarshal([]byte(msg.Data),&loginResMsg)
	if err != nil{
		return
	}
	//登录不成功
	if loginResMsg.Code != 200{
		err = errors.New(loginResMsg.Err)
		return
	}

	//登录成功，
	//开一个goroutine跟服务器保持连接
	go keepAlive(conn)

	//登录成功，修改并显示在线列表
	fmt.Println("当前在线用户id如下：")
	//显示在线userId(遍历loginResMsg.OnlineUserIds)
	for _,v := range loginResMsg.OnlineUserIds{
		fmt.Println("userId =",v)
		//修改在线用户列表map[int]*model.User
		userMgr.updateOnlineUser(v,message.USER_ONLINE)
	}
	//将自己也添加入在线用户列表
	userMgr.updateOnlineUser(userId,message.USER_ONLINE)

	//登录成功，设置自身信息
	curUser.UserId = userId
	curUser.UserState = message.USER_ONLINE
	curUser.Conn = conn
	
	//显示登录成功后的主菜单
	showLoginSucessMenu(userId)

	return
}

//向服务器发package，并接收服务器发送过来的package
func sendPkgAndRecievePkg(conn *net.Conn,data []byte)(message message.Message ,err error){
	
	//实例utils.transfer
	transfer := utils.Transfer{
		Conn : *conn,
	}
	//发送package writePkg()
	err = transfer.WritePkg(data)
	if err != nil{
		return
	}
	
	//接收服务器返回的Message
	message,err = transfer.ReadPkg()
	if err != nil{
		fmt.Println("userHandle.go  sendPkgAndRecievePkg() transfer.ReadPkg() err =",err)
		return
	}
	return
}
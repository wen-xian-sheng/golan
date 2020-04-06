package main
import (
	"fmt"
	"net"
	"go_code/chatroom/utils"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/handle"
	"io"
)

//等待客户端连接
func waitConn(listen net.Listener){
	for{
		//等待客户端连接
		conn,err := listen.Accept()
		//延时关闭conn
		defer conn.Close()
		if err != nil{
			fmt.Println("listen.Accept err =",err)
		}else{
			fmt.Println("有一个客户端接入")
			//连接成功,开启goroutine处理
			go process(conn)
		}
	}
} 

//响应客户端的请求
func process(conn net.Conn){
	//读取客户端发送的内容并返回message.Message
	transfer := utils.Transfer{
		Conn : conn,
	}
	for{//不停的循环读取该链接发送过来的消息
		msg,err := transfer.ReadPkg()
		if err != nil {
			if err == io.EOF{
				fmt.Println("客户端关闭连接，服务器也关闭连接")
				return
			}
			fmt.Println("processor.go readAndReturn(conn) err =",err)
			return
		}

		//switch Message.Type,对应处理
		switch msg.Type{
		case message.LoginMsgType:
			//用户登录
			userHandle := handle.UserHandle{
				Conn : conn,
			}
			err = userHandle.Login(&msg)
			if err != nil{
				fmt.Println("processor.go  process() Login() err =",err)
				break
			}
			fmt.Println("processor.go  process() 用户已登录 ",msg.Data)
		case message.RegistMsgType:
			//用户注册
			userHandle := handle.UserHandle{
				Conn : conn,
			}
			err = userHandle.Regist(&msg)
			if err != nil{
				fmt.Println("processor.go  process() Regist() err =",err)
				return
			}
			//return ,结束该链接，不在读取客户端发送的内容
			return  
			
		//接收到群发消息Type
		case message.SendGroupMsgType:
			fmt.Println("processor.go  process() 群发消息接收到的数据msg =",msg)
			var smsHandler *handle.SmsHandle
			err := smsHandler.SendGroupMsg(&msg)
			if err != nil{
				fmt.Println("processor.go process() 群发消息处理错误")
			}
		default:
			fmt.Println("processor.go  process() 客户端发送的请求没有相应的处理业务")
			return
		}
	}
}

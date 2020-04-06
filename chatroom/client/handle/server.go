package handle
import (
	"go_code/chatroom/utils"
	"fmt"
	"net"
	"go_code/chatroom/common/message"
	"encoding/json"
)

//跟服务器保持连接,获取服务器的消息
func keepAlive(conn net.Conn){
	defer conn.Close()
	//实例utils.transfer
	transfer := utils.Transfer{
		Conn : conn,
	}

	for{
		//接收服务器的消息
		// fmt.Println("server.go keepAlive() 等待接收消息")
		msg,err := transfer.ReadPkg()
		if err != nil {
			fmt.Println("handle/server.go keepAlive()接收服务器的消息 err =",err)
			return;
		}
		fmt.Println("handle/server.go keepAlive()接收服务器的消息 message =",msg)

		switch msg.Type {
		//如果是通知用户状态改变
		case message.NotifyStateMsgType: 
			//msg.Data反序列化成NotifyStateMsg
			var notifyStateMsg message.NotifyStateMsg
			err = json.Unmarshal([]byte(msg.Data),&notifyStateMsg)
			if err != nil{
				fmt.Println("handle/server.go keepAlive()反序列化异常 err =",err)
				return
			}
			//修改在线用户列表
			userMgr.updateOnlineUser(notifyStateMsg.UserId,notifyStateMsg.State)
			//显示在线列表
			userMgr.printOnlineUser()
		//接收到群发
		case message.SendGroupMsgType: 
			//打印出群发的消息
			smsMgr.printSendGroupMsg(&msg)
		default:
			fmt.Println("server.go keepAlive() 对服务器返回的message.Type没有相应处理方法")	
		}
		
	}
	
}

//显示登录成功后的主菜单
func showLoginSucessMenu(userId int){
	for{
		fmt.Printf("--------恭喜%d登录成功------------\n",userId)
		fmt.Printf("--------1.显示在线用户列表------------\n")
		fmt.Printf("--------2.发送消息------------\n")
		fmt.Printf("--------3.信息列表------------\n")
		fmt.Printf("--------4.退出系统------------\n")
		fmt.Printf("请选择(1-4)：")

		var key int
		fmt.Scanf("%d\n",&key)

		switch key{
			case 1 :
				userMgr.printOnlineUser()
			case 2 :
				fmt.Println("请输入要发送的消息：")
				var content string
				fmt.Scanln(&content)
				err := smsHandle.sendGroupMsg(content)
				if err != nil {
					fmt.Println("server.go showLoginSucessMenu 群发消息失败 err =",err)
				}
			case 3 :
				fmt.Println("选择3")
			case 4 :
				fmt.Println("选择4")
				// os.Exit(0)
				return
			default :
				fmt.Println("输入有误，请重新输入")
		}
	}

}
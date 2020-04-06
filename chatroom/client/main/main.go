package main
import (
	"fmt"
	"go_code/chatroom/client/handle"
)

//function() print show_main_view
func showMainView(){
	//实例UserHandle
	var userHandle handle.UserHandle
	//define key for recieving what user scan
	var key int 
	//define loop 判断是否循环
	var loop bool = true

	for loop {
		fmt.Println("--------------------------welcome chatroom system-------------------------")
		fmt.Println("\t\t\t 1.登录聊天室")
		fmt.Println("\t\t\t 2.注册账号")
		fmt.Println("\t\t\t 3.退出系统")
		fmt.Print("please choose 1-3: ")

		//recieve what user scan
		fmt.Scanln(&key)

		switch key {
		case 1:
			fmt.Println("登录中")
			err := userHandle.Login()
			if err != nil{
				fmt.Println("登录失败，err =",err)
			}
		case 2:
			fmt.Println("regist")
			err := userHandle.Regist()
			if err != nil{
				fmt.Println("注册失败，err =",err)
			}else{
				fmt.Println("注册成功")
			}
			// loop = false
		case 3:
			fmt.Println("login out")
			loop = false
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}
	
}

func main(){
	showMainView()
}
package main
import (
	"fmt"
	"net"
	"time"
	"go_code/chatroom/server/model"
)

var userDao *model.UserDao

//初始化
func init(){
	//初始化redis连接池
	initPool("localhost:6379",8,0,time.Second * 30)
	fmt.Println(pool)
	// 单例UserDao
	model.InitUserDao(pool)
}

func main(){
	//监听端口
	fmt.Println("server开始监听8889端口")
	listen,err := net.Listen("tcp","0.0.0.0:8889")
	defer listen.Close()
	//监听失败，返回
	if err != nil {
		fmt.Println("net.Listen err =",err)
		return
	}

	//等待客户端连接
	waitConn(listen)
}
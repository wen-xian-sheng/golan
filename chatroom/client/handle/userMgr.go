package handle
import (
	"go_code/chatroom/common/model"
	"go_code/chatroom/common/message"
	"fmt"
)

type UserMgr struct{
	users map[int]*model.User  //用户列表
	onLineUsers map[int]*model.User  //在线用户列表
}

var userMgr UserMgr

func init(){
	userMgr = UserMgr{
		//初始化好友
		users : make(map[int]*model.User,10),
		//make map 在线好友列表
		onLineUsers : make(map[int]*model.User,10),
	}
}

//修改在线用户列表map[int]*model.User
func(this *UserMgr) updateOnlineUser(userId int,state int)(err error){
	switch state {
	//如果状态是在线则新构建user再加入到map
	case message.USER_ONLINE:
		user := model.User{
			UserId : userId,
		}
		this.onLineUsers[userId] = &user
	} 
	
	return
}

//显示在线列表
func(this *UserMgr) printOnlineUser(){
	fmt.Println("当前在线列表更新如下：")
	for k,_ := range this.onLineUsers{
		fmt.Println("userId =",k)
	}
}

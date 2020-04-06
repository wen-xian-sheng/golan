package handle

import (
	"fmt"
)

type UserMgr struct{
	onlineUser map[int] *UserHandle
}

//整个服务器生命周期只有一个实例
var userMgr *UserMgr

//初始化userMgr
func init(){
	fmt.Println("userMgr.go init()")
	userMgr = &UserMgr{
		onlineUser : make(map[int] *UserHandle,1024),
	}
}

//增加在线用户
func(this *UserMgr) insertOnlineUser(userHandle *UserHandle){
	this.onlineUser[userHandle.UserId] = userHandle
}

//删除在线用户
func(this *UserMgr) deleteOnlineUser(userHandle *UserHandle){
	delete(this.onlineUser,userHandle.UserId)
}

//返回全部在线在线用户
func(this *UserMgr) getAllOnlineUser() (onlineUsers map[int]*UserHandle){
	return this.onlineUser
}

//返回单个在线用户
func(this *UserMgr) getOnlineUserByUserId(userId int)(userHandle *UserHandle,err error){
	var ok bool
	userHandle,ok = this.onlineUser[userId]
	if !ok {
		err = fmt.Errorf("该用户%d不在线",userId)
	}
	return 
}
package model
import (
	"github.com/garyburd/redigo/redis"
	"go_code/chatroom/common/model"
	"encoding/json"
	"sync"
	"fmt"
)

type UserDao struct{
	pool *redis.Pool
}

var  (
	MyUserDao *UserDao
)
//工厂模式，单例UserDao
var mutex sync.Mutex
func InitUserDao(pool *redis.Pool){
	
	// if &MyUserDao == nil{
	// 	mutex.Lock()
	// 	defer mutex.Unlock()
	// 	if MyUserDao == nil{
	// 		MyUserDao = &UserDao{
	// 			pool: pool,
	// 		}
	// 	}
	// }
	MyUserDao = &UserDao{
		pool: pool,
	}

}

//往redis插入user
func(this *UserDao) Insert(user *model.User)(err error) {
	//从redis连接此获取连接
	conn := this.pool.Get()
	defer conn.Close()
	
	// 根据userId从redis获取user
	_,err = redis.String(conn.Do("hget","users",user.UserId))

	//如果获取成功，则表示在redis已经存在，返回
	if err == nil{
		err = ERROR_USER_EXIT
		return
	}

	//获取不到则表示redis不存在该user，则插入到redis
	//将user序列化
	var data []byte
	data,err = json.Marshal(user)
	_,err = conn.Do("hset","users",user.UserId,data)
	if err != nil{
		return
	}

	return
}

//根据userId从redis获取user
func(this *UserDao) GetByUserId(userId int,userPwd string)(user *User,err error){
	//从redis连接此获取连接
	conn := this.pool.Get()
	defer conn.Close()
	var res string
	res,err = redis.String(conn.Do("hget","users",userId))
	if err != nil{
		if err == redis.ErrNil{
			//表示在hash中没有userId
			err = ERROR_USER_NOT_EXIT
		}
		return
	}
	//将user json字符串反序列化成user
	err = json.Unmarshal([]byte(res),&user)
	if err != nil {
		fmt.Println("userDao.go GetByUserId() 反序列化失败 res =",res)
		return
	}
	//将redis获取到的userPwd与用户输入的userPwd比较
	if user.UserPwd != userPwd{
		err  = ERROR_USERPWD
	}
	return
}
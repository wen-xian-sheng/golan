package main
import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var pool *redis.Pool
func initPool(address string,maxIdle int,maxActive int,idleTiemout time.Duration){
	pool = &redis.Pool{
		MaxIdle: maxIdle, //最大空闲连接数
		MaxActive: maxActive, //最大和数据库连接数,0表示没有数据
		IdleTimeout: idleTiemout, //最大空闲时间
		Dial: func()(redis.Conn, error){ //初始化连接的代码，连接哪个IP的redis
			return redis.Dial("tcp",address)
		},
	}
}
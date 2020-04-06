package utils
import (
	"net"
	"go_code/chatroom/common/message"
	"encoding/binary"
	"encoding/json"
)

type Transfer struct{
	Conn net.Conn
	buf [8096]byte
}

// 读取发送过来的package并返回message.Message
func(this *Transfer) ReadPkg()(msg message.Message,err error){
	for{
		//读取长度切片
		_,err = this.Conn.Read(this.buf[:4])
		//conn读取长度失败，返回
		if err != nil{
			return
		}
		//将长度切片转成uint32
		len := int(binary.BigEndian.Uint32(this.buf[:4]))

		//读取message切片
		var n int
		n,err = this.Conn.Read(this.buf[:len])
		//conn读取message失败，返回
		if n != len || err != nil{
			return
		}
		//将message切片反序列化
		err = json.Unmarshal(this.buf[:len],&msg)
		//message反序列失败，返回
		if err != nil{
			return
		}
		return
	}
}

//写入并发送package
func(this *Transfer) WritePkg(data []byte) (err error){
	//获取msg包的长度（[]byte）
	var msgLen uint32
	msgLen = uint32(len(data))
	
	//将长度（uint32）转成 切片
	binary.BigEndian.PutUint32(this.buf[0:4],msgLen)

	//发送msg包的长度
	var n int
	n,err = this.Conn.Write(this.buf[:4])
	if n != 4 || err != nil{
		return
	}

	//发送message.Message的[]byte
	n,err = this.Conn.Write(data)
	if n != int(msgLen) ||err !=nil{
		return
	}
	return
}

//构建message.Message并序列化,发送给客服端
func(this *Transfer) MarshalAndWritePkg(structData interface{},msgType string) (err error){

	var msg message.Message
	//序列化传入的struct
	var bytes []byte
	bytes,err = json.Marshal(structData)
	if err != nil{
		return
	}
	msg.Type = msgType
	msg.Data = string(bytes)
	bytes,err = json.Marshal(msg)
	if err != nil{
		return
	}
	//响应给客户端
	err = this.WritePkg(bytes)
	
	return
}

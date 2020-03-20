package model
import (
	"fmt"
)

type CustomerModel struct{
	Id int
	Name string
	Gender string
	Age int
	Email string
}

//使用工厂模式，返回一个实例
func NewCustomer(id int,name string,gender string,
	age int,email string) CustomerModel{
		customer := CustomerModel{
			Id : id,
			Name : name,
			Gender : gender,
			Age : age,
			Email : email,
		}
		return customer
}

//返回一个没有id的customer
func NewCustomer2(name string,gender string,
	age int,email string) CustomerModel{
		customer := CustomerModel{
			Name : name,
			Gender : gender,
			Age : age,
			Email : email,
		}
		return customer
}

//返回详细信息
func(this *CustomerModel) GetInfo() string{
	info := fmt.Sprintf("%v \t %v \t %v \t %v \t %v",this.Id,this.Name,this.Gender,this.Age,this.Email)
	return info
}
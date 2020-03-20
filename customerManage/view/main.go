package main
import (
	"fmt"
	"go_code/customerManage/service"
	"go_code/customerManage/model"
)

type customerView struct{
	key int //接收用户输入
	loop bool //结束循环
	customerService *service.CustomerService //customerService实例
}

//显示customer详细信息
func(this *customerView) list(){
	fmt.Println("----------------------------list customer info---------------------------")
	fmt.Println("id \t name \t gender  age \t email")
	customers := this.customerService.List()
	for i :=0;i < len(customers);i++{
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("-----------------------   list customer info over------------------------\n\n")
}

//添加customer
func(this *customerView) add(){
	fmt.Println("---------------------------添加客户--------------------------------------------")
	fmt.Println("请输入姓名")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("请输入性别")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("请输入年龄")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("请输入邮箱")
	email := ""
	fmt.Scanln(&email)

	//封装到customer
	customer := model.NewCustomer2(name,gender,age,email)
	//调用service的add方法添加到切片
	if this.customerService.Add(customer){
		fmt.Println("--------------------------添加成功-------------------------------------------")
	} else{
		fmt.Println("--------------------------添加失败-------------------------------------------")
	}
}

//删除customer
func(this *customerView) delete(){
	fmt.Println("请输入要删除的id（-1不删除）")
	id := -1
	fmt.Scanln(&id) //接收输入的id
	if id == -1{
		fmt.Println("退出删除")
		return
	}
	fmt.Println("你确定要删除",id,"Y/N")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "Y" || choice == "y" {
		//调用service 的Delete
		if this.customerService.Delete(id){
			fmt.Println("删除成功")
		} else {
			fmt.Println("删除失败")
		}
	}
	fmt.Println("删除失败")
}

//修改customer
func(this *customerView) update(){
	fmt.Println("请输入要修改的id")
	id := -1
	fmt.Scanln(&id)
	//根据id获取customer
	index := this.customerService.FindIndex(id)
	if index == -1 {
		fmt.Println("没有找到customer")
		return
	}
	customer := this.customerService.List()[index]

	//修改名字	
	fmt.Printf("姓名：%v 请输入修改后的名字(回车不修改)：",customer.Name)
	name := ""
	fmt.Scanln(&name)
	if name == "" {
		name = customer.Name
	}
	//修改性别
	fmt.Printf("性别：%v 请输入修改后的性别(回车不修改)：",customer.Gender)
	gender := ""
	fmt.Scanln(&gender)
	if gender == "" {
		gender = customer.Gender
	}
	//修改年龄
	fmt.Printf("年龄%v 请输入修改后的年龄：",customer.Age)
	age := 0
	fmt.Scanln(&age)
	//修改邮箱
	fmt.Printf("邮箱%v 请输入修改后的邮箱(回车不修改)\n",customer.Email)
	email := ""
	fmt.Scanln(&email)
	if email == "" {
		email = customer.Email
	}
	//构建new customer
	newCustomer := model.NewCustomer(id,name,gender,age,email)

	if this.customerService.Update(index,newCustomer) {
		fmt.Println("修改成功")
	} else {
		fmt.Println("修改失败")
	}
}

func(this *customerView) showView(){
	for{
		fmt.Println("------------------------customer manage system------------------------")
		fmt.Println("                         1.查看客户资料")
		fmt.Println("                         2.删除客户")
		fmt.Println("                         3.增加客户")
		fmt.Println("                         4.修改客户")
		fmt.Println("                         5.退出")
		fmt.Print("请选择")

		//key 接受用户输入
		fmt.Scanln(&this.key)
		//判断key
		switch this.key {
		case 1 :
			this.list()
		case 2 :
			this.delete()
		case 3 :
			this.add()
		case 4 :
			this.update()
		case 5 :
			fmt.Println("you exit system")
			this.loop = false
		default:
			fmt.Println("你输入有误，请重新输入")
		}
		if !this.loop {
			break
		}
	}
}

func main(){
	this := customerView{
		key : 0,
		loop : true,
	}
	//初始化customerService
	this.customerService = service.NewCustomerService()
	this.showView()
}
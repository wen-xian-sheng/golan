package service
import (
	"go_code/customerManage/model"
)
type CustomerService struct{
	//定义用户切片
	customers []model.CustomerModel
	//定义切片长度
	customerNum int
}

//返回customerService实例指针
func NewCustomerService() *CustomerService{
	customerService := &CustomerService{}
	customerService.customerNum = 1
	//初始化一个customer
	customer := model.NewCustomer(1,"张三","男",18,"zs@qq.com")
	//将顾客append进切片
	customerService.customers = append(customerService.customers,customer)
	//返回customerService地址
	return customerService
}

//返回customer切片
func(customerService *CustomerService) List() []model.CustomerModel{
	return customerService.customers
}

//增加customer
func(customerService *CustomerService) Add(customer model.CustomerModel) bool{
	//定义id
	customerService.customerNum++
	customer.Id = customerService.customerNum
	//将customer加入切片
	customerService.customers = append(customerService.customers,customer)
	return true
}

//删除customer
func(this *CustomerService) Delete(id int) bool{
	result := false
	//调用本地方法，查找index
	index := this.FindIndex(id)
	//如果下标不等于-1，则删除
	if index != -1 {
		this.customers = append(this.customers[:index],this.customers[index + 1:]...)
		result = true
	}
	return result

}
//跟据id查找customer所在切片的index
func(this *CustomerService) FindIndex(id int) int{
	index := -1
	//遍历切片
	for i := 0;i < len(this.customers);i++{
		if this.customers[i].Id == id{
			index = i
		}
	}
	return index
}

//修改customer信息
func(this *CustomerService) Update(index int,customer model.CustomerModel) bool {
	result := false 
	this.customers[index] = customer
	result = true
	return result
}
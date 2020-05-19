package encap_test

import (
	"fmt"
	"testing"
	"unsafe"
)

//go的面向对象编程

type Employee struct {
	Id   string
	Name string
	Age  int
}

//数据封装
func TestInitObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	e1 := Employee{Name: "Jone", Age: 21}
	e2 := new(Employee) //返回的是指针

	e2.Id = "2"
	e2.Name = "Rose"
	e2.Age = 30

	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2) //&{2 Rose 30},&表示取址，表示一个地址
	t.Logf("e is %T", e)
	t.Logf("e2 is %T", e2) //*encap_test.Employee，带*号，表示指针类型
}

//行为封装
//定义到实例上,实例e会被复制一份再调用该方法
func (e Employee) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name)) //与调用的实例地址不同
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

//定义到指针类型上,推荐第二种方式，减少值复制的开销
// func (e *Employee) String() string {
// 	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
// 	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
// }

func TestInitMethod(t *testing.T) {
	//e := &Employee{Id: "100", Name: "Gogo", Age: 5}
	e := Employee{Id: "100", Name: "Gogo", Age: 5}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}

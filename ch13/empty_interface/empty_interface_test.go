package empty_interface

import (
	"fmt"
	"testing"
)

//空接口可以接收任何类型
func DoSomething(p interface{}) {
	//断言方式判断类型
	// if i, ok := p.(int); ok {
	// 	fmt.Println("Integer", i)
	// 	return
	// }

	// if s, ok := p.(string); ok {
	// 	fmt.Println("string", s)
	// 	return
	// }

	// fmt.Println("Unknown Type")

	//switch方式
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("Unknown type")
	}
}

func TestEmptInterface(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
}

//倾向于小接口定义
//大接口可以由小接口组合而成
//方法尽量只依赖于功能最小的接口

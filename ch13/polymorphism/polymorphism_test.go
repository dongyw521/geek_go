package polymorphism

import (
	"fmt"
	"testing"
)

//多态

type Code string
type Programmer interface {
	WriteHelloworld() Code
}

// func(p * Programmer) WriteHelloworld() Code{
// 	return ""
// }

type GoProgrammer struct {
}

func (g GoProgrammer) WriteHelloworld() Code {
	return "fmt.Println(\"Hello World\")"
}

type JavaProgrammer struct {
}

func (j *JavaProgrammer) WriteHelloworld() Code {
	return "System.Out.Println(\"Hello World\")"
}

//传递指针，还是实例，取决于接口方法是定义在指针上，还是实例上；GoProgrammer定义到了实例上，JavaProgrammer定义到了指针上，一般
//采用定义到指针上
func WriteFirstProgram(p Programmer) {
	//fmt.Println(p.WriteHelloworld())
	fmt.Printf("%T,%v", p, p.WriteHelloworld())
}

func TestPolymorphism(t *testing.T) {
	g := GoProgrammer{}      //返回实例
	j := new(JavaProgrammer) //返回实例指针
	WriteFirstProgram(g)
	WriteFirstProgram(j)
}

//空接口和断言

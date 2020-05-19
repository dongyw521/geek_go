package interface_test

import "testing"

//定义接口
type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

//ducktype方式实现接口
func (g *GoProgrammer) WriteHelloWorld() string {
	return "Hello World!"
}

func TestInterface(t *testing.T) {
	var g Programmer
	g = new(GoProgrammer)
	t.Log(g.WriteHelloWorld())
}

//接口
//1.非侵入性，实现不依赖接口定义
//2.接口定义可以包含在使用者包内

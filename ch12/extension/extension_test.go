package extension

import (
	"fmt"
	"testing"
)

//复用，不能重写，父类变量也无法接收子类实例
type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Println("...")
}

func (p *Pet) SpeakTo(hoster string) {
	p.Speak()
	fmt.Println(hoster)
}

type Dog struct {
	//复用,可以通过p调用pet的方法,两个完全独立
	//p *Pet
	//命名嵌套类型
	Pet //无需重新定义方法，dog自动会有和pet一样的方法，但不能重写，两者之间还是独立的，父类变量也无法接收子类实例
}

// func (d *Dog) Speak() {
// 	d.p.Speak()
// 	fmt.Println("Wang!")
// }

// func (d *Dog) SpeakTo(hoster string) {
// 	d.Speak()
// 	d.p.SpeakTo(hoster)
// }

func TestDog(t *testing.T) {
	//var d Pet = new(Dog) 报错
	d := new(Dog)
	d.SpeakTo("Dong")
}

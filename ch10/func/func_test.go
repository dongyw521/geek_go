package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

//go语言里的所有参数传递都是值传递

//多个返回值函数
func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	//忽略某个返回值
	c, _ := returnMultiValues()
	t.Log(a, b)
	t.Log(c)
}

//包装函数inner，返回一个函数，该函数可以打印inner执行时间，并返回inner的执行结果
func TimeSpan(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("Time span:", time.Since(start).Seconds())
		return ret
	}
}

//模拟耗时函数
func SlowFunc(op int) int {
	time.Sleep(time.Second * 2)
	return op
}

func TestTimeSpent(t *testing.T) {
	fn := TimeSpan(SlowFunc)
	r := fn(20)
	t.Log(r)
}

//可变长参数,本质还是转换为数组
func Sum(ops ...int) int {
	ret := 0
	for _, n := range ops {
		ret += n
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
}

//defer 延迟调用执行函数
func Clear() {
	fmt.Println("Clear Resources")
}

func TestDefer(t *testing.T) {
	defer Clear() //延迟执行clear
	fmt.Println("start")
	panic("error")
	//fmt.Println("end") panic之后的代码不执行
}

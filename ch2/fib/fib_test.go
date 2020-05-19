package fib

import "testing"

//输出斐波那契数列
func TestFib(t *testing.T) {
	t.Log("Fib")
	a := 1
	b := 1
	t.Log(a)
	t.Log(b)
	for i := 0; i < 5; i++ {
		tmp := b    //tmp定义为存放上一次的计算结果，也就是离当前位置最近的前一个元素的值，并且在开始循环时，赋值为b，其实b才是存放的上一次的计算结果
		b = tmp + a //计算当前位置的元素的值，a存储的是当前位置上上个元素的值
		a = tmp     //赋值为tmp，即下一个要计算位置的上上个元素的值，为下一次计算做准备
		t.Log(b)
	}
}

//交换两个变量的值
func TestExchange(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a
	t.Log(a, b)
}

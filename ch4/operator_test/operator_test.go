package operator_test

import "testing"

//数组可以直接通过==比较，维数，元素，及元素顺序相同，两个数组相等
func TestCompareArr(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	c := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	t.Log(a == c)
}

//go按位清零运算符&^,右面为1，左边置位0;右边为0，左边保持不变
const (
	Readable   = 1 << iota //左移位运算0001
	Writable               //0010
	Executable             //0100
)

func TestBit(t *testing.T) {
	a := 7              //0111
	a = a &^ Readable   //0110
	a = a &^ Writable   //0100
	a = a &^ Executable //0000
	t.Log(a)
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}

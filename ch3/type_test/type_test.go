package type_test

import "testing"

//go不支持隐式类型转换
type MyInt int64 //类的别名也不支持隐式转换

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	//b = a//报错，不支持隐式类型转换，必须显示转换
	b = int64(a)
	t.Log(a, b)

	var c MyInt
	//c = b//类的别名也不支持隐式转换
	c = MyInt(b)
	t.Log(c)
}

//go不支持指针运算
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//aPtr = aPtr + 1//go不支持指针运算
	t.Log(a, aPtr)
	t.Logf("%T,%T", a, aPtr)
}

//string为值类型，初始值为空字符串nil，而不是null
func TestString(t *testing.T) {
	var a string
	t.Log("*" + a + "*")
	t.Log(len(a)) //输出0
}

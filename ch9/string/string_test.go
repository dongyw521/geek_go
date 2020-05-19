package string_test

import "testing"

//字符串本质是一个byte组成的slice切片,里面的值不可变
func TestString(t *testing.T) {
	var s string
	t.Log(s) //空字符串
	s = "hello"
	t.Log(s)
	t.Log(len(s)) //返回的是字符串中的字节个数

	s = "\xE4\xB8\xA5"
	s = "\xE4\xBB\xA5"

	t.Log(s)
	t.Log(len(s))

	//Unicode是字符的编码方式，Utf-8是unicode编码后的存储方式
	s = "中"
	t.Log(s)
	t.Log(len(s))  //输出slice中字节个数
	c := []rune(s) //获取s的Unicode编码
	t.Log(c)
	t.Log(len(c))
	t.Logf("中 unicode %x", c[0]) //16进制
	t.Logf("中 utf8 %x", s)

}

//字符串函数

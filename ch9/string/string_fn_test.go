package string_fn

import (
	"strconv"
	"strings"
	"testing"
)

//切割和连接函数
func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",") //切割
	t.Log(parts)
	for _, part := range parts {
		t.Log(part)
	}

	partJoin := strings.Join(parts, "-") //连接
	t.Log(partJoin)
}

//类型转换函数
func TestStringConv(t *testing.T) {
	s := strconv.Itoa(10) //int转换为字符串
	t.Log("string" + s)
	//字符串转整形需要有错误判断
	//i := strconv.Atoi("10") //字符串转int,这样写编译错误。
	if i, err := strconv.Atoi("10"); err == nil { //中间用分号;，不能用逗号
		t.Log(10 + i)
	}
	//t.Log(i)
}

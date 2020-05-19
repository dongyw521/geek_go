package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int //声明arr为数组类型
	t.Log(arr[0])

	arr1 := [3]int{1, 2, 3}
	t.Log(arr1)

	arr2 := [...]int{1, 3, 5, 6, 7} //动态长度
	t.Log(arr2, len(arr2))
}

func TestArrayTravel(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7}

	//for遍历方式
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}

	//go特有的遍历方式，使用range关键字，for第一个参数为索引，第二个为元素本身
	for idx, e := range arr {
		t.Log(idx, e)
	}

	for _, e := range arr { //忽略索引的遍历方式，需要一个下划线占位符，不能空着
		t.Log(e)
	}
}

//数组截取
func TestArraySection(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5}
	b := a[1:]  //索引为1开始的全部元素
	c := a[2:2] //为空，第二个参数也为索引，但不包括第二个
	d := a[2:3]
	e := a[:3]
	//f := a[-3:-1]//不支持负数
	t.Log(b)
	t.Log(c)
	t.Log(d) //输出3
	t.Log(e) //输出1,2,3
	//t.log(f)
}

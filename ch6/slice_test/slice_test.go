package slice_test

import "testing"

//切片，类似可变长数组
func TestSliceInit(t *testing.T) {
	//init方式1
	var s0 []int
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4, 5}
	t.Log(len(s1), cap(s1))

	//make方式
	s2 := make([]int, 3, 5)
	t.Log(s2[0], s2[1], s2[2])
	//t.Log(s2[0], s2[1], s2[2], s2[3]) //抛出异常，out of range
	t.Log(len(s2), cap(s2))
	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(len(s2), cap(s2))
}

//slice切片的cap会成倍数的增加，并且每次都需要把s0赋值为新的append后的切片，因为它是一个结构体，每次append后，地址会变。

func TestSliceGrowing(t *testing.T) {
	s0 := []int{}
	for i := 0; i < 10; i++ {
		s0 = append(s0, i)
		t.Log(len(s0), cap(s0))
	}
}

//slice是一个结构体，其中有一个字段指向了一块连续的存储空间
//slice共享存储空间
func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	//Q2和summer指向同一块存储空间，即year
	summer[0] = "Unknown" //Q2和year中的值都会变
	t.Log(Q2, len(Q2), cap(Q2))
	t.Log(year, len(year), cap(year))
}

//切片和切片不能比较，只能和nil空值比较，是否为空
//数组可以比较两个数组是否完全一样

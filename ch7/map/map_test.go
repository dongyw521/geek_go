package map_test

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[1])
	t.Logf("len of m1 is %d", len(m1))

	m2 := map[int]int{}
	m2[3] = 4
	t.Logf("len of m2 is %d", len(m2))

	m3 := make(map[int]int, 10)        //10为capacity,但无法通过cap(m3)获取到
	t.Logf("len of m3 is %d", len(m3)) //长度为0
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1]) //并不存在的key，值默认为0，这样通过值是否为空null无法判断该key是否存在
	m1[3] = 3
	//通过if判断是否存在key
	if v, ok := m1[3]; ok {
		t.Logf("key 3 is existed:%d", v)
	} else {
		t.Logf("key 3 is not existed")
	}
}

//map遍历,range方式
func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 3, 3: 4, 5: 6}
	for k, v := range m1 {
		t.Log(k, v)
	}
}

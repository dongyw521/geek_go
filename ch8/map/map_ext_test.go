package map_ext

import "testing"

//map里存放func
func TestMapWithFuncValue(t *testing.T) {
	m1 := map[int]func(op int) int{}
	m1[1] = func(op int) int { return op }
	m1[2] = func(op int) int { return op * op }
	m1[3] = func(op int) int { return op * op * op }
	t.Log(m1[1](2), m1[2](2), m1[3](2))
}

//实现set
func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 1
	if mySet[n] {
		t.Logf("%d is existing.", n)
	} else {
		t.Logf("%d is not existing.", n)
	}

	mySet[3] = true
	delete(mySet, 1)
	//n = 3
	if mySet[n] {
		t.Logf("%d is existing.", n)
	} else {
		t.Logf("%d is not existing.", n)
	}
}

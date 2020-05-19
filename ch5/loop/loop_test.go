package loop_test

import "testing"

//go 循环只有一个关键字for
func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 { //相当于while循环
		t.Log(n)
		n++
	}
}

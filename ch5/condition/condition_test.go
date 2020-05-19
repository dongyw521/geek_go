package condition_test

import "testing"

func TestMultiSel(t *testing.T) {
	//if支持多段写法，中间用分号，用于函数多返回值，比较方便
	if a := 1 == 1; a {
		t.Log("1==1")
	}
}

//switch在go里不需要break;
func TestSwitch1(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2: //一个case后面多个条件
			t.Log("Even") //每个case不需要手动break
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("Unknown")
		}
	}
}

//switch 当if条件使用，case后面不是常量，而是条件
func TestSwitch2(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch { //switch后面也没有变量
		case i%2 == 0: //一个case后面不是常量而是条件
			t.Log("Even2") //每个case不需要手动break
		case i%2 == 1:
			t.Log("Odd2")
		default:
			t.Log("Unknown2")
		}
	}
}

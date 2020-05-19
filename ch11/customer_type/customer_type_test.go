package customer_type

import (
	"fmt"
	"testing"
	"time"
)

//自定义类型
type IntCov func(op int) int //类似一个委托

//包装函数inner，返回一个函数，该函数可以打印inner执行时间，并返回inner的执行结果
func TimeSpan(inner IntCov) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("Time span:", time.Since(start).Seconds())
		return ret
	}
}

//模拟耗时函数
func SlowFunc(op int) int {
	time.Sleep(time.Second * 2)
	return op
}

func TestTimeSpent(t *testing.T) {
	fn := TimeSpan(SlowFunc)
	r := fn(20)
	t.Log(r)
}

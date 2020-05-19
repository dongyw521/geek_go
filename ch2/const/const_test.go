package const_test

import "testing"

func TestConst(t *testing.T) {
	const (
		Monday = 1 + iota
		Tuesday
		Wednesday
	)

	const (
		Readable   = 1 << iota //左移位运算001
		Writable               //010
		Executable             //100
	)
	t.Log(Monday, Tuesday, Wednesday)
	t.Log(Readable, Writable, Executable)
}

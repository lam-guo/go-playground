package pointer

import (
	"testing"
	"unsafe"
)

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// aPtr = aPtr + 1 //不支持指针运算
	t.Log(a, aPtr)
	// uintptr支持指针运算
	b := uintptr(unsafe.Pointer(&a)) + 1
	t.Log(a, aPtr, b, &b)
	t.Logf("%T %T", a, aPtr)
}

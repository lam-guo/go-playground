package main

import "C"

import "fmt"

// 用go实现SayHello也是可以的（c++,其他语言都可以）
//
//export SayHello
func SayHello(s *C.char) {
	fmt.Print(C.GoString(s))
}

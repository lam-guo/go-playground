package main

//#include <hello.h>
import "C"

func main() {
	println("hello cgo")
	C.SayHello(C.CString("Hello, World\n"))
}

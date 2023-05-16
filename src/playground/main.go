package main

func foo() {
	println("foo")
}


// 在1.20.4的版本会输出 1，2，3
// 在1.20.3的版本会输出 foo,foo,foo
func main() {
	fn := foo
	for _, fn = range list {
		fn()
	}
}

var list = []func(){
	func() {
		println("1")
	},
	func() {
		println("2")
	},
	func() {
		println("3")
	},
}

package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {

	switch v := p.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("unknow")
	}
}

func TestEmptyInterface(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
	DoSomething(true)
}

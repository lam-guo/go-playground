package playground

import (
	"fmt"
	"testing"
)

type A interface {
	get
	Set()
}

type get interface {
	Get() bool
}

type a_impl struct {
}

func (a *a_impl) Get() bool {
	fmt.Println("get")
	return false
}

func (a *a_impl) Set() {
	fmt.Print("set")
}

func NewA() A {
	return &a_impl{}
}

func Newget() get {
	return &a_impl{}
}

func TestNew(t *testing.T) {
	a := NewA()
	g := Newget()
	a.Get()
	g.Get()
}

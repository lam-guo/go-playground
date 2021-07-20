package extension

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	Pet
}

func (d *Dog) Speak() {
	fmt.Print("wang!")
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.Speak()
}

type Code string
type Programmer interface {
	Write() Code
}
type GoP struct {
}

func (p *GoP) Write() Code {
	return "fmt.Println(hello)"
}

type JavaP struct{}

func (j *JavaP) Write() Code {
	return "system.out.Println(hello)"
}

func writFP(p Programmer) {
	fmt.Printf("%T %v\n", p, p.Write())

}

func TestPoly(t *testing.T) {
	goPro := new(GoP)
	javaPro := new(JavaP)
	writFP(goPro)
	writFP(javaPro)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))
	fmt.Println(time.Now())
	runtime.GC()
	fmt.Println(time.Now())
}

//继承,重写...
package main

import "fmt"

type Base struct {
	Name string
}

func (base *Base) Foo() {
	fmt.Println("Base foo func... ")
}

func (base *Base) Bar() {
	fmt.Println("Base bar func...")
}

type Foo struct {
	Base
	sex string
}

func (foo *Foo) Bar() {
	foo.Base.Bar() //重写BAR，先调用基类BAR
	fmt.Println("Foo bar func....")
}

func main() {
	b := new(Base)
	b.Bar()
	b.Foo()

	f := new(Foo)
	f.Bar()
	f.Base.Bar()
	f.Base.Foo()
}

package main

import (
	"fmt"
)

func main() {
	var a [20]int
	fmt.Println(a)

	//索引19的元素设置为1，其它为0
	b := [20]int{19: 1}
	fmt.Println(b)

	var c = [...]int{1, 2, 3, 4, 5}
	fmt.Println(c)

	var d = [...]int{0: 1, 1: 2, 3: 4, 5, 6}
	fmt.Println(d)

	e := [...]int{9: 1}
	var p *[10]int = &e //指向数组的指针
	fmt.Println(p)

	x, y := 1, 2
	f := [...]*int{&x, &y} //指针数组
	fmt.Println(f)

	p1 := new([10]int) //new得到的是一个指向数组的指针
	fmt.Println(p1)

	a1 := [10]int{}
	a1[1] = 2
	fmt.Println(a1)
	p2 := new([10]int)
	p2[1] = 2
	fmt.Println(p2)

	//2个[3]int 数组
	a3 := [2][3]int{
		{1, 1, 1},
		{2, 2, 2}}
	fmt.Println(a3)
}

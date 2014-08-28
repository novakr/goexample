package main

import (
	"fmt"
)

const (
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
)

func main() {
	fmt.Println("hello world 你好")

	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)

	a := 1
	var p *int = &a
	fmt.Println(p)

	if a, b := 1, 2; a < b {
		a = b
		fmt.Println(a)
		fmt.Println(b)
	}
	fa := 1
	for {
		fa++
		if fa > 5 {
			break
		}
		fmt.Println(fa)
	}

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	fmt.Println("##################")
	for a < 3 {
		a++
		fmt.Println(a)

	}

}

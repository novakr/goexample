package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)

	for i := 0; i < 10; i++ {
		go GO(c, i)
	}
	<-c

	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(a)
}

func GO(c chan bool, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	if index == 9 {
		c <- true
	}
}

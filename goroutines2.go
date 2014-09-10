package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go GO(c, i)
	}

	for i := 0; i < 10; i++ {
		<-c
	}

	ch2 := fenfa(100)
	sum := 0
	for i := 0; i < 10; i++ {
		sum = sum + <-ch2
	}
	fmt.Println("fenfa....", sum)

}

func GO(c chan bool, index int) {

	a := 0
	for i := 0; i <= 100000000; i++ {

		a += i
	}

	fmt.Println(index, a)

	//if index == 9 {
	c <- true
	//}

}
func fenfa(d int) (ch chan int) {
	c := make(chan int, 10)
	step := d / 10
	start := 0
	end := 0
	for i := 1; i <= 10; i++ {
		start = (i - 1) * step
		end = start + step
		if end > d {
			end = d
		}
		go Calc(start, end, c)
	}
	return c
}

func Calc(start int, end int, result chan int) {

	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}
	result <- sum
	fmt.Println(start, end, sum)
}

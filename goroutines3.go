package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go GO(&wg, i)
	}
	wg.Wait()

	var sum = 0
	for i := 1; i <= 50; i++ {
		sum = sum + i
	}
	fmt.Println(sum * 50)
}

func GO(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	wg.Done()

}

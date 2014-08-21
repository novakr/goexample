package main

import (
	"fmt"
	"runtime"
)

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	fmt.Println("CPU Num:", runtime.NumCPU())
}

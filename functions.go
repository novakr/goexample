package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func vals(a int, b int) (int, int) {
	return a + b, a - b
}

func main() {
	res := plus(1, 2)
	fmt.Println("1 + 2 = ", res)

	r1, r2 := vals(3, 5)
	fmt.Println("a + b = ", r1)
	fmt.Println("a - b = ", r2)

	r3, _ := vals(66, 5)
	fmt.Println("a + b = ", r3)

	sum(1, 3, 4, 9, 43, 23)

	nums := []int{1, 3, 4, 5, 6}
	sum(nums...)

}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

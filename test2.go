package main

import "fmt"

func SunAndProduct(A, B int) (int, int) {
	return A + B, A * B
}
func main() {
	var tmap map[string]int
	tmap = make(map[string]int, 3)
	fmt.Printf("len:%d\n", len(tmap))

	tmap["1"] = 1
	tmap["2"] = 2
	fmt.Printf("len:%d\n", len(tmap))

	var aSlice []
	x := 3
	y := 4
	xplusy, xtimesy := SunAndProduct(x, y)
	fmt.Printf("%d + %d = %d \n", x, y, xplusy)

	fmt.Printf("%d * %d = %d \n", x, y, xtimesy)

	sum := 0
	for index := 0; index < 10000; index++ {
		sum += index
	}
	fmt.Println("sum is equal to ", sum)

	sum1 := 1
	for sum1 < 10000 {
		sum1 += sum1
		if sum1 > 1000 {
			break
		}
	}
	fmt.Println("sum1 is equal to ", sum1)
}

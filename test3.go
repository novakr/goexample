package main

import "fmt"

func modify(array [10]int) {
	array[0] = 10
	fmt.Println("in modify() array values:", array)
}

func main() {
	array := [10]int{1, 2, 3, 4, 5}
	modify(array)
	fmt.Println("in main,array values:", array)

	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var mySlice []int = myArray[:5]

	for _, v := range myArray {
		fmt.Print(v, " ")
	}
	fmt.Println("")
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}

	fmt.Println("len(mySlice):", len(mySlice))
	fmt.Println("cap(mySlice):", cap(mySlice))

	mySlice = append(mySlice, 1, 2, 3, 4, 5, 6)
	fmt.Println("len(mySlice):", len(mySlice))
	fmt.Println("cap(mySlice):", cap(mySlice))
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	fmt.Println("88888888888888888888888888888888")
	var mySlice2 []int
	mySlice2 = make([]int, 5, 10)
	mySlice2 = []int{8, 9, 10}
	//mySlice2 := []int{8, 9, 10}
	mySlice = append(mySlice, mySlice2...)
	fmt.Println("len(mySlice):", len(mySlice))
	fmt.Println("cap(mySlice):", cap(mySlice))
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	mySlice3 := []int{11, 12}
	fmt.Println("####################")
	for _, v := range mySlice3 {
		fmt.Print(v, " ")
	}
	fmt.Println("##")
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 9}
	//copy(slice2, slice1)
	for _, v := range slice2 {
		fmt.Println(v, " ")
	}
	fmt.Println("####################")
	copy(slice1, slice2)
	for _, v := range slice1 {
		fmt.Print(v, " ")
	}
}

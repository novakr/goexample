//Copyright zengjr
//for example
//

/*
	for test
*/
package main

import (
	"fmt"
)

//
//main method
//
//
func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	//BUG(zengjr): #1 : for test....
}

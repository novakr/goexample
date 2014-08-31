package main

import (
	"fmt"
	//"time"
)

/*
func main() {
	go Go()
	time.Sleep(2 * time.Second)

	c := make(chan bool)
	go func() {
		fmt.Println("wa1 wa wa")
		c <- true
	}()
	<-c
}

func Go() {
	fmt.Println("wa wa wa")
}
*/
func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("wa wa")
		c <- true
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}

package main

import (
	"fmt"
	"time"
)

func main() {
	time1 := time.NewTimer(time.Second * 2)

	<-time1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

}

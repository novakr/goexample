package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args
	argswithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argswithoutProg)
	fmt.Println(arg)
}

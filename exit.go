package main

import (
	"fmt"
	"os"
)

//使用exit defer内容不会输出
func main() {
	defer fmt.Println("!")

	os.Exit(2)
}

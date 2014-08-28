//任意类型不定参数
//
package main

import "fmt"

func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value")
		case string:
			fmt.Println(arg, "is an string value")
		case int64:
			fmt.Println(arg, "is an int64 value")
		default:
			fmt.Println(arg, "is an unknown type")
		}
	}
}

func main() {
	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.234
	MyPrintf(v1, v2, v3, v4)

	var j int = 5

	a := func() func() {
		var i int = 10
		return func() {
			fmt.Printf("i,j:%d,%d\n", i, j)
		}
	}()
	a()
	j *= 2
	a()
}

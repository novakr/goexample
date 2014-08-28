package main

import "fmt"

func main() {
	const (
		i      = 100
		pi     = 3.1415926
		prefix = "GO_"
	)
	var tt map[string]int
	tt = make(map[string]int)
	tt["JAN"] = 100
	tt["SEN"] = 2

	s := "hello "
	m := `测试
			换行`
	a := s + m
	fmt.Printf("%s\n", a)
	fmt.Printf("aaa:", i)
	fmt.Println(pi, prefix)
	fmt.Println(tt["JAN"])
	delete(tt, "SEN")
	fmt.Println("长度:", len(tt))
	res, ok := tt["JAN"]
	if ok {
		fmt.Println("ok", res)
		if res > 1 {
			fmt.Println("hello")
		}
	} else {
		fmt.Println("error")
	}
	myFunc()
}

//test goto
func myFunc() {
	i := 0

HERE:

	fmt.Println(i)
	i++
	if i < 10 {
		goto HERE
	}

}

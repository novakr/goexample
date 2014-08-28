package main

import "fmt"

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)

	personDB["12345"] = PersonInfo{"12345", "tom", "ROOM203"}
	personDB["1"] = PersonInfo{"1", "JACK", "ROOM202"}

	person, ok := personDB["1234"]
	fmt.Println(len(personDB))
	if ok {
		fmt.Println("FOUND PERSON", person.Name, "with ID 1234")
	} else {
		fmt.Println("DID NOT FOUND PERSON WITH ID 1234")
	}
	for key, val := range personDB {
		fmt.Println("key...", key, "Val :", val)
	}
	val, flag := personDB["1"]
	if flag {
		fmt.Println(val, "。。。。。", flag)
	} else {
		fmt.Println(val, "eeeee", flag)
	}
	cc(2)
	switchfunc(11)
	var t = myfunc(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(t)
}
func myfunc(args ...int) int {
	var total = 0
	for _, v := range args {
		total = total + v
	}
	return total
}
func switchfunc(x int) {
	switch x {
	case 0:
		fallthrough
	case 1:
		fmt.Println("this is 1")
	case 2:
		fmt.Println("this is 2")
	case 3, 4, 5:
		fmt.Println("this is 3,4,5")
	default:
		fmt.Println("x is x")
	}
}

//test return 
//当函数有返回值时，不能在if else里面return 不然报错，说你没有返回值
func cc(x int) int {
	var temp int
	if x > 0 {
		temp = 1
	} else {
		temp = 0
	}
	return temp
}

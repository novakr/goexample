package main

import "fmt"

func main() {
	a := 1
	switch a {
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a = 1")
	default:
		fmt.Println("None")
	}
	fmt.Println("#################")
	switch {
	case a >= 0:
		fmt.Println("a>=0")
		fallthrough
	case a >= 1:
		fmt.Println("a>=1")
	default:
		fmt.Println("None")
	}
	fmt.Println("#################")
	switch a := 1; {
	case a >= 0:
		fmt.Println("a>=0")
		fallthrough
	case a >= 1:
		fmt.Println("a>=1")
	default:
		fmt.Println("None")
	}

LABEL1:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				break LABEL1
			}
		}
	}
	fmt.Println("ok")

LABEL2:
	for i := 1; i < 10; i++ {
		for {
			continue LABEL2
		}
	}
	fmt.Println("okQ")
}

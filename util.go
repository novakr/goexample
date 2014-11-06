package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	s1 := make([]string, 3)
	s1[0] = "a"
	s1[1] = "b"
	s1[2] = "c"

	var s2 = [2]int{1, 2}

	b, _ := Contains("a", s1)
	fmt.Println(b)
	b1, _ := Contains(2, s2)
	fmt.Println(b1)

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	b2, _ := Contains("k1", m)
	fmt.Println(b2)
}

//判断obj是否在target中,target支持类型为array,slice,map
func Contains(obj interface{}, target interface{}) (bool, error) {

	targetVal := reflect.ValueOf(target)

	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetVal.Len(); i++ {
			if targetVal.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetVal.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	}
	return false, errors.New("not in array")
}

package main

import "fmt"
import (
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
	User
	title string
}

func (u User) Hello() {
	fmt.Println("Hello World", u.Name)
}

func (u User) Hello1(name string) {
	fmt.Println("Hello ", name, ",my name is ", u.Name)
}
func main() {
	u := User{1, "ok", 12}
	Info(u)

	m := Manager{User: User{1, "0k", 12}, title: "123"}
	t := reflect.TypeOf(m)

	fmt.Printf("%#v\n", t.Field(0))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1}))

	x := 123
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(99999)
	fmt.Println(x)

	u1 := User{1, "ok", 12}
	Set(&u1)
	fmt.Println(u1)

	v1 := reflect.ValueOf(u1)
	mv := v1.MethodByName("Hello1")

	args := []reflect.Value{reflect.ValueOf("joe")}
	mv.Call(args)
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("xxxx")
		return
	} else {
		v = v.Elem()
	}

	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("ADB")
		return
	}

	if f := v.FieldByName("Name"); f.Kind() == reflect.String {
		f.SetString("bybyb")
	}
}
func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("xx")
		return
	}

	v := reflect.ValueOf(o)
	fmt.Println("Fields:")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s:%v=%v\n", f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%s:%v\n", m.Name, m.Type)
	}
}

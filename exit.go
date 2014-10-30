// Copyright © 2014 jr.zengjr@gmail.com All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file

/*
 包说明
*/
package main

import (
	"fmt"
	"os"
)

//函数说明
//os.Exit使用例子
//使用exit defer内容不会输出
func main() {
	defer fmt.Println("!")

	os.Exit(2)
}

// BUG(zengjr): #1: bug问题描述
// BUG(zengjr): #2: bug问题描述

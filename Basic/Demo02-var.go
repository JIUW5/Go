package main

import (
	"fmt"
)

// 函数外为全局变量
// 延迟声明
var name string
var age int

// 立即声明
var x = 1
var y = 3.5

var a int32 = 1
var b float32 = 1.2

// 变量块，归纳统一作用的变量，有助于代码阅读性
var (
	//推进：声明格式
	c = int64(1)
	d = float32(1.2)
	z = "xu"
)

func main() {
	name = "Allen xu"
	age = 10
	fmt.Printf("name = %s \nage = %d \n", name, age)

	fmt.Printf("x = %d , y = %f , a = %d , b = %f", x, y, a, b)
	fmt.Printf("c = %d , d = %f , z = %s\n", c, d, z)
	//调用函数
	add()
}

// 声明函数
func add() {
	//局部变量
	//短声明
	e := 1
	f := 2
	fmt.Printf("add函数：")
	fmt.Printf("\ne=%d f=%d , e+f=%d", e, f, e+f)

	//批量声明
	var g, h int
	g = 1
	h = 2
	fmt.Printf("\ng=%d h=%d g+h=%d", g, h, g+h)

	var (
		i = int64(1)
		j = int64(2)
	)
	fmt.Printf("\ni=%d j=%d i+j=%d", i, j, i+j)

}

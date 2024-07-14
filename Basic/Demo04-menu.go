package main

import (
	"fmt"
	"time"
)

// 无类型枚举定义
const (
	// A 使用iota关键字，声明此变量块里都是枚举类型
	A = iota
	B
	C
	D
)

// Color 自定义数据类型
type Color int

const (
	// Red 使用iota关键字，声明此变量块里都是枚举类型，且为Color数据类型
	Red Color = iota
	Blue
	Yellow
	Pink
)

func main() {
	fmt.Println(A, B, C, D)
	var y float32 = 1
	fmt.Println(A+y, B+y, C+y, D+y)
	fmt.Println(Red, Blue, Yellow, Pink)

	//使用SDK的枚举类型
	one := time.Monday

	fmt.Println(one)
}

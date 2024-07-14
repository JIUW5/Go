package main

import "fmt"

// NewInt 自定义类型
type NewInt int

// 声明常量
const (
	NAME  = 10
	COLOR = "Red"
)

// Y 使用自定义类型声明常量（开发的时候，可以直接通过变量类型得知此变量的作用）
const Y NewInt = 10

func main() {
	fmt.Printf("Y = [%d]", Y)
}

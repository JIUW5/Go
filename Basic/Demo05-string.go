package main

import "fmt"

func main() {
	name := "Allen"
	//字节数组
	nameCopy := []byte(name)
	nameCopy[0] = 'B'
	fmt.Println(name, string(nameCopy))

	var school string = "1"
	fmt.Println("school: ", school)

	//字符串的拼接
	var newName = name + " Niu" + " Bi"
	fmt.Println(newName)
	newName += "!"
	fmt.Println(newName)

	//字符串的判断
	if name == newName {
		fmt.Println("ok")
	} else {
		fmt.Println("no")
	}

	if name <= newName {
		fmt.Println("ok")
	} else {
		fmt.Println("no")
	}

	//多行字符串
	text := `Hi
Xu`
	fmt.Println(text)

}

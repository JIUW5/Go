package io

import (
	"bytes"
	"os"
	"testing"
)

func TestWriterReader(t *testing.T) {
	//bytes.Buffer：是一个实现了读写方法的可变大小的字节缓冲
	var b bytes.Buffer
	b.WriteString("!!! ")

	//os.Open：打开一个文件，返回一个文件对象
	file, err := os.Open("io.txt")
	if err != nil {
		t.Error(err)
	}
	//b.Bytes：返回缓冲的底层字节切片的拷贝,作用：将缓冲的内容转换为字节切片
	bs := b.Bytes()
	//string(bs)：将字节切片转换为字符串
	t.Log(string(bs))

	//b.string：返回缓冲的底层字节切片的拷贝,作用：将缓冲的内容转换为字符串 和 string(b.Bytes()) 一样
	s := b.String()
	t.Log(s)

	//b.Reset：重置缓冲，清空缓冲的内容
	b.Reset()

	//b.WriteString：将字符串s写入缓冲
	b.WriteString("ke")

	//os.OpenFile：打开一个文件，返回一个文件对象,os.O_RDWR：读写模式打开文件,0是什么意思：表示文件的权限，0表示没有权限
	file, err = os.OpenFile("io.txt", os.O_RDWR, 0)
	if err != nil {
		t.Error(err)
	}
	//权限为0怎么会写入文件？
	//因为，os.OpenFile打开文件的时候，是以os.O_RDWR模式打开文件的，所以，可以写入文件
	//b.WriteTo：将缓冲的内容写入文件
	_, err = b.WriteTo(file)
	if err != nil {
		t.Error(err)
	}

}

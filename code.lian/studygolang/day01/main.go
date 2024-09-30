package main

import "fmt"

const pi = 3.1415 //常量，不可修改，声明的时候进行赋值

const (

	//批量声明常量
	q = 100
	w = 200
	e //支持简写，相当于e = 200
)

//iota枚举关键字 与const搭配使用，const批量声明变量的时候，每出现一次iota，iota的值都会重置为0，出现iota的时候，每多一行，iota累加1

const (
	zz = iota
	xx
	vv
)

// 定义两个返回值的函数
func foo() (string, int) {
	return "lgh", 9000
}

func main() {

	var name string    // 标准声明
	var age int32 = 90 //  直接声明的时候赋值

	// 批量声明变量
	var (
		a int
		b string
		c bool
	)
	name = "jack112" //声明变量后赋值
	a = 10
	b = "变量"
	c = true
	fmt.Printf("%s的年龄是%d\n", name, age)
	fmt.Println(a, b, c)
	// 简短变量声明（只能在函数内部使用）
	lgh := "连光辉213" //  与  var lgh ="连光辉"  相同
	fmt.Println(lgh)

	//调用foo函数
	aa, bb := foo()
	fmt.Println(aa, bb)

	cc, _ := foo() // _ 用来接受不需要的值
	fmt.Println(cc)

	//不能重复申请同名变量
	//变量的赋值要写在方法里面
	//注意事项:1.函数外的每个语句都必须以关键字开始(var、func等)，:=不能使用在函数外。 2. _多用于占位，表示忽略值。

	fmt.Println("pi的值是", pi)
	fmt.Println(q, w, e)
	fmt.Println(zz, xx, vv)

	s3 := "hello中国" //utf-8 编码下一个中文占3个字节 len默认按照统计字节数量进行
	fmt.Println("s3字符的长度是", len(s3))

	for i := 1; i < len(s3); i++ {
		fmt.Printf("s3的第%d的位置是%d\n", i, s3[i])
	}

	//遇到中文的时候可以使用for range循环
	for k, v := range s3 {
		fmt.Printf("第%d的位置是%c\n", k+1, v)
	}

	// 字符串反转操作
	//方法1
	str1 := "hello"
	fmt.Printf("s1反转之前是%s\n", str1)

	byteArray := []byte(str1)
	str2 := ""
	for i := len(byteArray) - 1; i >= 0; i-- {
		str2 += string(byteArray[i])
	}
	fmt.Printf("s1反转之后是%s\n", str2)

	//方法2

	length := len(byteArray)
	for i := 0; i < length/2; i++ {
		byteArray[i], byteArray[length-i-1] = byteArray[length-i-1], byteArray[i]
	}
	fmt.Printf("第二种方法反转后，变为%s", byteArray)
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//非中文字符串截取
	s1 := "qwerqtyqu"
	s1 = string([]byte(s1)[1:3])

	//中文字符串截取
	s2 := "你好2345sadfg啊"
	s2 = string([]rune(s2)[1:3])
	fmt.Println(s1)
	fmt.Println(s2)

	//字符替换
	s4 := "1234qweqwewqewAWETXZCVrwqewr"
	s5 := strings.Replace(s4, "q", "bb", -1)
	s6 := strings.Replace(s4, "q", "bb", 3)
	fmt.Println("s5===>", s5)
	fmt.Printf("s6是%s\n", s6)

	//字符转换为大写
	s7 := strings.ToUpper(s4)
	fmt.Println("s7 ====>", s7)

	//字符转换为小写
	s8 := strings.ToLower(s7)
	fmt.Printf("s8的内容是%s\n", s8)

	//int+str 将字符的ASCII值相加
	c := 10 + 'a'
	fmt.Printf("c的值是：%c\n", c)

	//类型的强制转换   强制转换的时候转换的是变量值的类型，并不是变量的类型
	var p int32 = 100
	var n1 = float64(p)
	fmt.Println(n1)
	fmt.Printf("n1的type-->%T\n", n1)

	//其他数据类型转换为string---方式1:使用fmt.Sprintf
	var num1 int = 99
	var num2 float64 = 99.0999
	var b bool = false
	var myChar byte = 'h'
	var str string

	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str的类型是%T,str = %q\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str的类型是%T,str = %q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str的类型是%T,str = %q\n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str的类型是%T,str = %q\n", str, str)

	//其他数据类型转换为string---方式2:使用strconv中的Format函数

	var num3 int = 996
	var num4 float64 = 989.0999
	var b2 bool = false

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str的类型是%T,str = %q\n", str, str)

	str = strconv.FormatFloat(num4, 'f', 10, 64) // f表示格式 10表示小数点后保留10位，64表示float64
	fmt.Printf("str的类型是%T,str = %q\n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str的类型是%T,str = %q\n", str, str)

	//string转换为其他基础类型 使用strconv包
	//需要确保string类型能够转成有效的数据，例如不能把hello转换为整数10，golang会转换为0

	var num5 = "1234567"
	var num6 = "123.23456"
	var numRet int64
	var numRetF float64

	//转换为int
	numRet, _ = strconv.ParseInt(num5, 10, 64) //10是10进制，64是int64
	fmt.Printf("numRet的类型是%T,numRet= %v\n", numRet, numRet)

	//转换为float
	numRetF, _ = strconv.ParseFloat(num6, 64)
	fmt.Printf("numRet的类型是%T,numRet= %v\n", numRetF, numRetF)

}

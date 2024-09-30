package main

import (
	"fmt"
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

}

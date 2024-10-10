package main

import "fmt"

func main() {
	//var age byte
	//fmt.Println("请输入age:")
	//fmt.Scanln(&age)
	//if age >= 18 {
	//	fmt.Println("成年人")
	//} else if age <= 0 {
	//	fmt.Println("不合法")
	//} else {
	//	fmt.Println("未成年")
	//}

	//test1
	//var a, b int32
	//fmt.Println("请输入a:")
	//fmt.Scanln(&a)
	//fmt.Println("请输入b:")
	//fmt.Scanln(&b)
	//sum := a + b
	//if sum > 50 {
	//	fmt.Println("hello")
	//} else {
	//	fmt.Println("no")
	//}

	//test2
	//var a, b float64
	//fmt.Println("请输入a:")
	//fmt.Scanln(&a)
	//fmt.Println("请输入b:")
	//fmt.Scanln(&b)
	//if a > 10.0 && b > 20.0 {
	//	fmt.Println("sum:", a+b)
	//}

	//test3
	//var a, b int32
	//fmt.Println("请输入a:")
	//fmt.Scanln(&a)
	//fmt.Println("请输入b:")
	//fmt.Scanln(&b)
	//i3 := (a + b) % 3
	//i5 := (a + b) % 5
	//if i3 == 0 && i5 == 0 {
	//	fmt.Println("都能整除")
	//}

	//test4:判断闰年
	//var year int32
	//fmt.Println("请输入年份：")
	//fmt.Scanln(&year)
	//if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
	//	fmt.Printf("%v是闰年", year)
	//}

	//test5
	//var a int32
	//fmt.Println("请输入a:")
	//fmt.Scanln(&a)
	//if a == 100 {
	//	fmt.Println("100分")
	//} else if a <= 90 && a > 80 {
	//	fmt.Println("80<a<=90")
	//} else if a >= 60 && a <= 80 {
	//	fmt.Println("60=<a<=80")
	//} else {
	//	fmt.Println("none")
	//}

	//test6
	//var s float32
	//var gender int
	//fmt.Println("请输入时间")
	//fmt.Scanln(&s)
	//
	//if s <= 8.0 {
	//	fmt.Println("请输入性别")
	//	fmt.Scanln(&gender)
	//	if gender == 1 {
	//		fmt.Println("男子")
	//	} else {
	//		fmt.Println("女子")
	//	}
	//} else {
	//	fmt.Println("淘汰了")
	//}

	//test7
	//var m int32
	//fmt.Println("输入月份")
	//fmt.Scanln(&m)
	//var age int32
	//fmt.Println("输入年龄")
	//fmt.Scanln(&age)
	//if m <= 10 && m >= 4 {
	//	if age > 60 {
	//		fmt.Println("老人60")
	//	} else if age >= 18 {
	//		fmt.Println("成人20")
	//	} else {
	//		fmt.Println("儿童30")
	//	}
	//} else {
	//	if age >= 18 && age <= 60 {
	//		fmt.Println("成人40")
	//	} else {
	//		fmt.Println("票价20")
	//	}
	//}

	//test8
	//var weekday string
	//fmt.Println("请输入a-g：")
	//fmt.Scanln(&weekday)
	//switch weekday {
	//case "a":
	//	fmt.Println("周一")
	//case "b":
	//	fmt.Println("周二")
	//case "c":
	//	fmt.Println("周三")
	//case "d":
	//	fmt.Println("周四")
	//case "e":
	//	fmt.Println("周五")
	//case "f":
	//	fmt.Println("周六")
	//case "g":
	//	fmt.Println("周日")
	//default:
	//	fmt.Println("不合法")
	//}

	//输入a,b,c,d转换为大写，其他的输出other char
	//var useIn string
	//fmt.Println("请输入小写字母：")
	//fmt.Scanln(&useIn)
	//switch useIn {
	//case "a":
	//	fmt.Println("A")
	//case "b":
	//	fmt.Println("B")
	//case "c":
	//	fmt.Println("C")
	//case "d":
	//	fmt.Println("D")
	//default:
	//	fmt.Println("other char")
	//}
	//
	////score>=60 pass,  score<60 fail. other:不合法
	//
	//var score int32
	//fmt.Println("请输入score：")
	//fmt.Scanln(&score)
	//switch {
	//case score >= 60 && score <= 100:
	//
	//	fmt.Println("pass")
	//case score >= 0 && score < 60:
	//	fmt.Println("fail")
	//default:
	//	fmt.Println("不合法")
	//}
	//
	////根据输入的月份判断季节
	//
	//var month int32
	//fmt.Println("请输入month：")
	//fmt.Scanln(&month)
	//switch {
	//case month >= 1 && month <= 3:
	//	fmt.Println("春")
	//case month >= 4 && month <= 6:
	//	fmt.Println("夏")
	//case month >= 7 && month <= 9:
	//	fmt.Println("秋")
	//case month >= 10 && month <= 12:
	//	fmt.Println("冬")
	//default:
	//	fmt.Println("不合法")
	//}

	var str byte
	fmt.Println("请输入：")
	fmt.Scanf("%c", &str)
	switch {
	//case str >= 97 && str <= 101:
	//	fmt.Printf("%c", str-32)
	//default:
	//	fmt.Println("不合法")
	case 'a' <= str && str <= 'e':
		fmt.Printf("%c", str-32)
	default:
		fmt.Println("不合法啊")
	}

}

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
	var m int32
	fmt.Println("输入月份")
	fmt.Scanln(&m)
	var age int32
	fmt.Println("输入年龄")
	fmt.Scanln(&age)
	if m <= 10 && m >= 4 {
		if age > 60 {
			fmt.Println("老人60")
		} else if age >= 18 {
			fmt.Println("成人20")
		} else {
			fmt.Println("儿童30")
		}
	} else {
		if age >= 18 && age <= 60 {
			fmt.Println("成人40")
		} else {
			fmt.Println("票价20")
		}
	}
}
